package services

import (
	"context"
	"errors"
	"ticpin-backend/config"
	"ticpin-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	PassPrice          = 999.0
	PassDurationMonths = 3
)

var defaultBenefits = models.PassBenefits{
	TurfBookings:         models.BenefitCounter{Total: 2, Used: 0, Remaining: 2},
	DiningVouchers:       models.DiningVoucherBenefit{Total: 2, Used: 0, Remaining: 2, ValueEach: 250},
	EventsDiscountActive: true,
}

func GetActivePassByUserID(userID string) (*models.TicpinPass, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	col := config.GetDB().Collection("ticpin_passes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var pass models.TicpinPass
	if err := col.FindOne(ctx, bson.M{"user_id": objID, "status": "active"}).Decode(&pass); err != nil {
		return nil, err
	}
	return &pass, nil
}

func ApplyPass(userID, paymentID string, details models.TicpinPass) (*models.TicpinPass, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	col := config.GetDB().Collection("ticpin_passes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existing models.TicpinPass
	existsErr := col.FindOne(ctx, bson.M{"user_id": objID, "status": "active"}).Decode(&existing)
	if existsErr == nil {
		return nil, errors.New("active pass already exists")
	}

	profile, _ := GetProfileByUserID(userID)
	if profile != nil {
		if details.Name == "" {
			details.Name = profile.Name
		}
		if details.Phone == "" {
			details.Phone = profile.Phone
		}
		if details.Address == "" {
			details.Address = profile.Address
		}
		if details.Country == "" {
			details.Country = profile.Country
		}
		if details.State == "" {
			details.State = profile.State
		}
		if details.District == "" {
			details.District = profile.District
		}
	}

	now := time.Now()
	pass := &models.TicpinPass{
		ID:        primitive.NewObjectID(),
		UserID:    objID,
		Name:      details.Name,
		Phone:     details.Phone,
		Address:   details.Address,
		Country:   details.Country,
		State:     details.State,
		District:  details.District,
		PaymentID: paymentID,
		Price:     PassPrice,
		Status:    "active",
		StartDate: now,
		EndDate:   now.AddDate(0, PassDurationMonths, 0),
		Benefits:  defaultBenefits,
		Renewals:  []models.RenewalRecord{},
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err = col.InsertOne(ctx, pass)
	return pass, err
}

func RenewPass(passID, paymentID string) (*models.TicpinPass, error) {
	objID, err := primitive.ObjectIDFromHex(passID)
	if err != nil {
		return nil, err
	}

	col := config.GetDB().Collection("ticpin_passes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var pass models.TicpinPass
	if err := col.FindOne(ctx, bson.M{"_id": objID}).Decode(&pass); err != nil {
		return nil, errors.New("pass not found")
	}

	now := time.Now()
	newStart := pass.EndDate
	if now.After(pass.EndDate) {
		newStart = now
	}
	newEnd := newStart.AddDate(0, PassDurationMonths, 0)

	renewalRecord := models.RenewalRecord{
		RenewedAt: now,
		StartDate: newStart,
		EndDate:   newEnd,
		PaymentID: paymentID,
		Price:     PassPrice,
	}

	update := bson.M{
		"$set": bson.M{
			"status":     "active",
			"start_date": newStart,
			"end_date":   newEnd,
			"benefits":   defaultBenefits,
			"payment_id": paymentID,
			"updatedAt":  now,
		},
		"$push": bson.M{"renewals": renewalRecord},
	}

	if _, err := col.UpdateOne(ctx, bson.M{"_id": objID}, update); err != nil {
		return nil, err
	}

	pass.Status = "active"
	pass.StartDate = newStart
	pass.EndDate = newEnd
	pass.Benefits = defaultBenefits
	pass.PaymentID = paymentID
	pass.Renewals = append(pass.Renewals, renewalRecord)
	return &pass, nil
}

func UseTurfBooking(passID string) (*models.TicpinPass, error) {
	return useBenefit(passID, "turf")
}

func UseDiningVoucher(passID string) (*models.TicpinPass, error) {
	return useBenefit(passID, "dining")
}

func useBenefit(passID, benefitType string) (*models.TicpinPass, error) {
	objID, err := primitive.ObjectIDFromHex(passID)
	if err != nil {
		return nil, err
	}

	col := config.GetDB().Collection("ticpin_passes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var pass models.TicpinPass
	if err := col.FindOne(ctx, bson.M{"_id": objID, "status": "active"}).Decode(&pass); err != nil {
		return nil, errors.New("active pass not found")
	}

	var updateFields bson.M
	if benefitType == "turf" {
		if pass.Benefits.TurfBookings.Remaining <= 0 {
			return nil, errors.New("no turf bookings remaining")
		}
		pass.Benefits.TurfBookings.Used++
		pass.Benefits.TurfBookings.Remaining--
		updateFields = bson.M{
			"benefits.turf_bookings.used":      pass.Benefits.TurfBookings.Used,
			"benefits.turf_bookings.remaining": pass.Benefits.TurfBookings.Remaining,
			"updatedAt":                        time.Now(),
		}
	} else {
		if pass.Benefits.DiningVouchers.Remaining <= 0 {
			return nil, errors.New("no dining vouchers remaining")
		}
		pass.Benefits.DiningVouchers.Used++
		pass.Benefits.DiningVouchers.Remaining--
		updateFields = bson.M{
			"benefits.dining_vouchers.used":      pass.Benefits.DiningVouchers.Used,
			"benefits.dining_vouchers.remaining": pass.Benefits.DiningVouchers.Remaining,
			"updatedAt":                          time.Now(),
		}
	}

	if _, err = col.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": updateFields}); err != nil {
		return nil, err
	}
	return &pass, nil
}

func ExpireOldPasses() error {
	col := config.GetDB().Collection("ticpin_passes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.UpdateMany(ctx,
		bson.M{"status": "active", "end_date": bson.M{"$lt": time.Now()}},
		bson.M{"$set": bson.M{"status": "expired", "updatedAt": time.Now()}},
	)
	return err
}
