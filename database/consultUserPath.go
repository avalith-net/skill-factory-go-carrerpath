package database

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultUserPath(userId string, pathId string) (*models.RelatadPath, bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("careerpath")
	col := db.Collection("userPath")

	condition := bson.M{
		"userid":     userId,
		"userpathid": pathId,
	}

	var userPath models.RelatadPath
	err := col.FindOne(ctx, condition).Decode(&userPath)
	idRelation := userPath.RelationID.Hex()
	if err != nil {
		return nil, false, "", err
	}
	return &userPath, true, idRelation, nil
}
