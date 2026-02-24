package services

import (
	"context"
	"errors"
	"ticpin-backend/config"
	"ticpin-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func LoginOrCreateOrganizer(email, password string) (*models.Organizer, bool, error) {
	collection := config.GetDB().Collection("organizers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var org models.Organizer
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&org)

	if err != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, false, err
		}

		org = models.Organizer{
			ID:                primitive.NewObjectID(),
			Email:             email,
			Password:          string(hashed),
			OrganizerCategory: []string{},
			IsVerified:        false,
			CreatedAt:         time.Now(),
		}

		if _, err := collection.InsertOne(ctx, org); err != nil {
			return nil, false, err
		}

		return &org, true, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(org.Password), []byte(password)); err != nil {
		return nil, false, errors.New("invalid password")
	}

	return &org, false, nil
}

func SendPlayOrganizerOTP(email string) error {
	collection := config.GetDB().Collection("organizers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	otp := config.GenerateOTP()
	expiry := time.Now().Add(10 * time.Minute)

	_, err := collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{
		"$set": bson.M{"otp": otp, "otpExpiry": expiry},
	})
	if err != nil {
		return err
	}

	return config.SendPlayOTP(email, otp)
}

func VerifyOrganizerOTP(email, otp string) (*models.Organizer, error) {
	collection := config.GetDB().Collection("organizers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var org models.Organizer
	if err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&org); err != nil {
		return nil, errors.New("organizer not found")
	}

	if org.OTP != otp {
		return nil, errors.New("invalid otp")
	}

	if time.Now().After(org.OTPExpiry) {
		return nil, errors.New("otp expired")
	}

	_, err := collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{
		"$set": bson.M{"isVerified": true, "otp": "", "otpExpiry": time.Time{}},
	})
	if err != nil {
		return nil, err
	}

	org.IsVerified = true
	return &org, nil
}

func GetOrganizerByID(id string) (*models.Organizer, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := config.GetDB().Collection("organizers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var org models.Organizer
	if err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&org); err != nil {
		return nil, err
	}

	return &org, nil
}

func CreateOrganizerProfile(profile *models.OrganizerProfile) error {
	profile.ID = primitive.NewObjectID()
	profile.CreatedAt = time.Now()
	profile.UpdatedAt = time.Now()

	collection := config.GetDB().Collection("organizer_profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, profile)
	return err
}

func GetOrganizerProfileByID(organizerID string) (*models.OrganizerProfile, error) {
	objID, err := primitive.ObjectIDFromHex(organizerID)
	if err != nil {
		return nil, err
	}

	collection := config.GetDB().Collection("organizer_profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var profile models.OrganizerProfile
	if err := collection.FindOne(ctx, bson.M{"organizerId": objID}).Decode(&profile); err != nil {
		return nil, err
	}

	return &profile, nil
}

func UpdateOrganizerProfile(organizerID string, profile *models.OrganizerProfile) error {
	objID, err := primitive.ObjectIDFromHex(organizerID)
	if err != nil {
		return err
	}

	profile.UpdatedAt = time.Now()

	collection := config.GetDB().Collection("organizer_profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.UpdateOne(ctx, bson.M{"organizerId": objID}, bson.M{"$set": profile})
	return err
}
