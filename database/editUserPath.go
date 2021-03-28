package database

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EditUserPath(pathUser models.RelatadPath, originalPath models.Path, relationID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("careerpath")
	colUsers := db.Collection("relation")

	pathTemplate := make(map[string]interface{})

	fmt.Println(pathUser.Path)
	fmt.Println(pathUser.Path.Description)
	fmt.Println(pathUser.Path.TechnicalSkills)
	fmt.Println(pathUser.Path.SoftSkills)
	emptyPath := models.RelatadPath{}

	if pathUser.Path.Description != originalPath.Description {
		pathUser.Path.Description = originalPath.Description
	}

	if !reflect.DeepEqual(pathUser.Path, emptyPath) {
		originalPath.TechnicalSkills = append(originalPath.TechnicalSkills, pathUser.Path.TechnicalSkills...)
		pathTemplate["path"] = originalPath.TechnicalSkills
	}

	// if len(pathUser.Path.TechnicalSkills) > 0 {
	// 	pathTemplate["technicalSkills"] = pathUser.Path.TechnicalSkills
	// }
	// if len(pathUser.Path.SoftSkills) > 0 {
	// 	pathTemplate["softSkills"] = pathUser.Path.SoftSkills
	// }

	updateToString := bson.M{
		"$set": pathTemplate,
	}
	fmt.Println(updateToString)

	objectID, _ := primitive.ObjectIDFromHex(relationID)
	filter := bson.M{"_id": bson.M{"$eq": objectID}}

	// opts := options.Update().SetUpsert(true)

	result, err := colUsers.UpdateOne(ctx, filter, updateToString)
	if err != nil {
		return false, err
	}
	fmt.Println(result)
	return true, nil
}
