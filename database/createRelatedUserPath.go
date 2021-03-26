package database

import (
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"golang.org/x/net/context"
	"time"
)

func CreateRelatedUserPath(rel models.RelatadPath) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("careerpath")
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, rel)

	if err != nil {
		return false, err
	}
	return true, nil
}
