package database

import (
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"golang.org/x/net/context"
)

func CreateRelatedUserPath(rel models.RelatadPath) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("careerpath")
	col := db.Collection("userPath")

	_, err := col.InsertOne(ctx, rel)

	if err != nil {
		return false, err
	}
	return true, nil
}
