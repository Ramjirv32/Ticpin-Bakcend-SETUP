package services

import (
	"context"
	"io"
	"ticpin-backend/config"
	"ticpin-backend/models"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProfile(profile *models.Profile) error {
	profile.ID = primitive.NewObjectID()
	profile.CreatedAt = time.Now()
	profile.UpdatedAt = time.Now()

	collection := config.GetDB().Collection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, profile)
	return err
}

func GetProfileByUserID(userID string) (*models.Profile, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	collection := config.GetDB().Collection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var profile models.Profile
	err = collection.FindOne(ctx, bson.M{"userId": objID}).Decode(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func UpdateProfile(userID string, profile *models.Profile) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	profile.UpdatedAt = time.Now()

	collection := config.GetDB().Collection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.UpdateOne(ctx, bson.M{"userId": objID}, bson.M{"$set": profile})
	return err
}

func UploadProfilePhoto(file io.Reader, userID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := config.GetCloudinary().Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:   "ticpin/profiles",
		PublicID: userID,
	})
	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}

func UpdateProfilePhoto(userID, photoURL string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	collection := config.GetDB().Collection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.UpdateOne(ctx, bson.M{"userId": objID}, bson.M{"$set": bson.M{"profilePhoto": photoURL, "updatedAt": time.Now()}})
	return err
}
