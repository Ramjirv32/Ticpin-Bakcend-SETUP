package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Offer struct {
	Discount string `bson:"discount" json:"discount"`
	Code     string `bson:"code" json:"code"`
}

type Dining struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrganizerID  primitive.ObjectID `bson:"organizer_id" json:"organizer_id"`
	Name         string             `bson:"name" json:"name"`
	Location     string             `bson:"location" json:"location"`
	Rating       float64            `bson:"rating" json:"rating"`
	IsOpen       bool               `bson:"is_open" json:"is_open"`
	OpenTime     string             `bson:"open_time" json:"open_time"`
	Phone        string             `bson:"phone" json:"phone"`
	BannerURL    string             `bson:"banner_url" json:"banner_url"`
	GalleryURLs  []string           `bson:"gallery_urls" json:"gallery_urls"`
	MenuURLs     []string           `bson:"menu_urls" json:"menu_urls"`
	Facilities   []string           `bson:"facilities" json:"facilities"`
	VenueName    string             `bson:"venue_name" json:"venue_name"`
	VenueAddress string             `bson:"venue_address" json:"venue_address"`
	Offers       []Offer            `bson:"offers" json:"offers"`
	FAQs         []FAQ              `bson:"faqs" json:"faqs"`
	Status       string             `bson:"status" json:"status"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
