package database

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
)

//PathUpdate is a function
func CreatePath(pathUser models.Path) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("careerpath")
	col := db.Collection("path")

	_, err := col.InsertOne(ctx, pathUser)
	if err != nil {
		return false, err
	}
	return true, nil
}
