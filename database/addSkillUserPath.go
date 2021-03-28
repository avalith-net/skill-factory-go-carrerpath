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

func AddSkillUserPath(userPath models.RelatadPath, originalPath *models.RelatadPath, relationID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("careerpath")
	collUsers := db.Collection("userPath")

	pathTemplate := make(map[string]interface{})

	emptyPath := models.RelatadPath{}

	if userPath.Path.Description != originalPath.Path.Description {
		userPath.Path.Description = originalPath.Path.Description
	}

	if !reflect.DeepEqual(userPath, emptyPath) {
		originalPath.Path.TechnicalSkills = append(originalPath.Path.TechnicalSkills, userPath.Path.TechnicalSkills...)
		originalPath.Path.SoftSkills = append(originalPath.Path.SoftSkills, userPath.Path.SoftSkills...)
		pathTemplate["path"] = originalPath.Path
	}

	updateToString := bson.M{
		"$set": pathTemplate,
	}
	fmt.Println(updateToString)

	objectID, _ := primitive.ObjectIDFromHex(relationID)
	filter := bson.M{"_id": bson.M{"$eq": objectID}}

	result, err := collUsers.UpdateOne(ctx, filter, updateToString)
	if err != nil {
		return false, err
	}
	fmt.Println(result)
	return true, nil
}
