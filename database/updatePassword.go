package database

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PasswordUpdate(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("careerpath")
	col := db.Collection("user")

	registro := make(map[string]interface{})

	user.Password, _ = EncryptPassword(user.Password)

	if len(user.Password) > 0 {
		registro["password"] = user.Password
	}

	if len(user.Email) > 0 {
		registro["email"] = user.Email
	}

	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}
	return true, nil

}
