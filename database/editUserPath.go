package database

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EditUserPath(pathUser models.Path, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("careerpath")
	col := db.Collection("path")

	pathTemplate := make(map[string]interface{})

	if len(pathUser.Description) > 0{
		pathTemplate["description"] = pathUser.Description
	}
	if len(pathUser.TechnicalSkills) > 0 {
		pathTemplate["technicalSkills"] = pathUser.TechnicalSkills
	}
	if len(pathUser.SoftSkills) > 0 {
		pathTemplate["softSkills"] = pathUser.SoftSkills //body en json
	}

	updateToString := bson.M{
		"$set": pathTemplate,
	}

	objectID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objectID}}

	_, err := col.UpdateOne(ctx, filter, updateToString)
	if err != nil {
		return false, err
	}
	return true, nil
}
