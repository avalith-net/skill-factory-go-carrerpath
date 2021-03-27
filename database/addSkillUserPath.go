package database

import (
	"context"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddSkillUserPath(path models.RelatadPath) (bool,error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("careerpath")
	col := db.Collection("testupdateskill")

	_, err := col.InsertOne(ctx, path)

	if err != nil {
		return false, err
	}
	return true, nil

}
