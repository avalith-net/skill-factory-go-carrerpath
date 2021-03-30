package database

import (
	"context"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func GetNotifications(id string)  (models.Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("careerpath")
	col := db.Collection("notification")

	var notification models.Notification

	err, _ := col.Find(ctx, bson.M{})

	if err != nil {
		return notification, err
	}

	return notification, nil
}
