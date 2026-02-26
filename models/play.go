package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Play struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrganizerID  primitive.ObjectID `bson:"organizer_id" json:"organizer_id"`
	TurfName     string             `bson:"turf_name" json:"turf_name"`
	About        string             `bson:"about" json:"about"`
	PlayOptions  []string           `bson:"play_options" json:"play_options"`
	Location     string             `bson:"location" json:"location"`
	VenueName    string             `bson:"venue_name" json:"venue_name"`
	VenueAddress string             `bson:"venue_address" json:"venue_address"`
	Duration     string             `bson:"duration" json:"duration"`
	BannerURL    string             `bson:"banner_url" json:"banner_url"`
	GalleryURLs  []string           `bson:"gallery_urls" json:"gallery_urls"`
	FAQs         []FAQ              `bson:"faqs" json:"faqs"`
	Terms        string             `bson:"terms" json:"terms"`
	Status       string             `bson:"status" json:"status"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
