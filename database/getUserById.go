package database

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetUserById receives an id and returns the user for that id
func GetUserById(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("careerpath")
	col := db.Collection("user")
	objID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": objID}
	var user models.User
	err := col.FindOne(ctx, condition).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
