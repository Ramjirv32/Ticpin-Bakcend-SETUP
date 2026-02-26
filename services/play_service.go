package services

import (
	"context"
	"ticpin-backend/config"
	"ticpin-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePlay(p *models.Play) error {
	p.ID = primitive.NewObjectID()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	col := config.GetDB().Collection("plays")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := col.InsertOne(ctx, p)
	return err
}

func GetAllPlays() ([]models.Play, error) {
	col := config.GetDB().Collection("plays")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var plays []models.Play
	if err := cursor.All(ctx, &plays); err != nil {
		return nil, err
	}
	return plays, nil
}

func GetPlayByID(id string) (*models.Play, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	col := config.GetDB().Collection("plays")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var p models.Play
	if err := col.FindOne(ctx, bson.M{"_id": objID}).Decode(&p); err != nil {
		return nil, err
	}
	return &p, nil
}
