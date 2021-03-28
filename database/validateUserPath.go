package database

import (
	"context"
	"reflect"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateUserPath(userPath models.RelatadPath, relationID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("careerpath")
	col := db.Collection("userPath")

	pathTemplate := make(map[string]interface{})

	emptyPath := models.RelatadPath{}

	if !reflect.DeepEqual(userPath, emptyPath) {
		pathTemplate["path"] = userPath.Path
	}
	updateToString := bson.M{
		"$set": pathTemplate,
	}

	objectID, _ := primitive.ObjectIDFromHex(relationID)
	filter := bson.M{"_id": bson.M{"$eq": objectID}}

	_, err := col.UpdateOne(ctx, filter, updateToString)
	if err != nil {
		return false, err
	}
	return true, nil

}
