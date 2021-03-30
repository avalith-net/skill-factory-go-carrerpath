package database

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetNotifications(page int64) ([]*models.Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("careerpath")
	col := db.Collection("notification")

	var notifications []*models.Notification

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	cursor, err := col.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var notification models.Notification
		err := cursor.Decode(&notification)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}

	err = cursor.Err()
	if err != nil {
		return nil, err
	}
	cursor.Close(ctx)
	return notifications, nil
}
