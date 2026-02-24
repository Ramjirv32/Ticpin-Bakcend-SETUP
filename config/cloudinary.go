package config

import (
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

var CloudinaryClient *cloudinary.Cloudinary

func InitCloudinary() error {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return err
	}

	CloudinaryClient = cld
	return nil
}

func GetCloudinary() *cloudinary.Cloudinary {
	return CloudinaryClient
}
