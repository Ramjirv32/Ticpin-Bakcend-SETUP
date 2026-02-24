package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organizer struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name              string             `bson:"name" json:"name"`
	Email             string             `bson:"email" json:"email"`
	Password          string             `bson:"password" json:"-"`
	OrganizerCategory []string           `bson:"organizerCategory" json:"organizerCategory"`
	OTP               string             `bson:"otp" json:"-"`
	OTPExpiry         time.Time          `bson:"otpExpiry" json:"-"`
	IsVerified        bool               `bson:"isVerified" json:"isVerified"`
	CreatedAt         time.Time          `bson:"createdAt" json:"createdAt"`
}

type OrganizerProfile struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrganizerID       primitive.ObjectID `bson:"organizerId" json:"organizerId"`
	Name              string             `bson:"name" json:"name"`
	Email             string             `bson:"email" json:"email"`
	OrganizerCategory []string           `bson:"organizerCategory" json:"organizerCategory"`
	Address           string             `bson:"address" json:"address"`
	Country           string             `bson:"country" json:"country"`
	State             string             `bson:"state" json:"state"`
	District          string             `bson:"district" json:"district"`
	ProfilePhoto      string             `bson:"profilePhoto" json:"profilePhoto"`
	CreatedAt         time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt         time.Time          `bson:"updatedAt" json:"updatedAt"`
}
