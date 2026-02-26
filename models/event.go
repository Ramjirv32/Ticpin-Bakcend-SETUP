package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FAQ struct {
	Question string `bson:"question" json:"question"`
	Answer   string `bson:"answer" json:"answer"`
}

type Event struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrganizerID      primitive.ObjectID `bson:"organizer_id" json:"organizer_id"`
	Name             string             `bson:"name" json:"name"`
	Category         string             `bson:"category" json:"category"`
	Date             time.Time          `bson:"date" json:"date"`
	Time             string             `bson:"time" json:"time"`
	Location         string             `bson:"location" json:"location"`
	VenueName        string             `bson:"venue_name" json:"venue_name"`
	VenueAddress     string             `bson:"venue_address" json:"venue_address"`
	BannerURL        string             `bson:"banner_url" json:"banner_url"`
	GalleryURLs      []string           `bson:"gallery_urls" json:"gallery_urls"`
	About            string             `bson:"about" json:"about"`
	Language         string             `bson:"language" json:"language"`
	Duration         string             `bson:"duration" json:"duration"`
	TicketsNeededFor string             `bson:"tickets_needed_for" json:"tickets_needed_for"`
	ArtistName       string             `bson:"artist_name" json:"artist_name"`
	ArtistImageURL   string             `bson:"artist_image_url" json:"artist_image_url"`
	PriceStartsFrom  float64            `bson:"price_starts_from" json:"price_starts_from"`
	FAQs             []FAQ              `bson:"faqs" json:"faqs"`
	Terms            string             `bson:"terms" json:"terms"`
	Status           string             `bson:"status" json:"status"`
	CreatedAt        time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt        time.Time          `bson:"updatedAt" json:"updatedAt"`
}
