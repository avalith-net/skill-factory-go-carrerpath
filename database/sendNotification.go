package database

import (
	"context"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"time"
)

func SendNotification(mNotification models.Notification) error{
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("careerpath")
	colNotification := db.Collection("notification")

	_, err := colNotification.InsertOne(ctx, mNotification)
	if err != nil {
		return err
	}
	return nil
}
