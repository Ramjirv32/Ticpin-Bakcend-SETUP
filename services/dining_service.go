package services

import (
	"context"
	"ticpin-backend/config"
	"ticpin-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDining(d *models.Dining) error {
	d.ID = primitive.NewObjectID()
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()

	col := config.GetDB().Collection("dinings")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := col.InsertOne(ctx, d)
	return err
}

func GetAllDinings() ([]models.Dining, error) {
	col := config.GetDB().Collection("dinings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var dinings []models.Dining
	if err := cursor.All(ctx, &dinings); err != nil {
		return nil, err
	}
	return dinings, nil
}

func GetDiningByID(id string) (*models.Dining, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	col := config.GetDB().Collection("dinings")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var d models.Dining
	if err := col.FindOne(ctx, bson.M{"_id": objID}).Decode(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
