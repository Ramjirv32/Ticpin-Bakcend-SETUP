package services

import (
	"context"
	"ticpin-backend/config"
	"ticpin-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user *models.User) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()

	collection := config.GetDB().Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	return err
}

func LoginUser(phone string) (*models.User, error) {
	collection := config.GetDB().Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByID(id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := config.GetDB().Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
