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

	register := make(map[string]interface{})

	//tomodPath.TechnicalSkills = append(tomodPath.TechnicalSkills, pathUser.TechnicalSkills...)

	if len(pathUser.TechnicalSkills) > 0 {
		register["technicalSkills"] = pathUser.TechnicalSkills
	}
	//tomodPath.SoftSkills = append(tomodPath.SoftSkills, pathUser.SoftSkills...)
	if len(pathUser.SoftSkills) > 0 {
		register["softSkills"] = pathUser.SoftSkills
	}

	updtString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateMany(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}
