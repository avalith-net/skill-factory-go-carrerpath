package database

import (
	"context"
	"fmt"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	Path pathInterface = &path{}
)

type path struct{}

type pathInterface interface {
	GetPathByID(string) (models.Path, error)
}

func (user *path) GetPathByID(ID string) (models.Path, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // releases resources if slowOperation completes before timeout elapses
	db := MongoCN.Database("careerpath")
	col := db.Collection("path")

	var path models.Path
	// ObjectIDFromHex creates a new ObjectID from a hex string. It returns an error if the hex string is not a
	// valid ObjectID.
	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id": objID,
	}
	if err := col.FindOne(ctx, condition).Decode(&path); err != nil {
		fmt.Println("path not found " + err.Error())
		return path, err
	}
	return path, nil
}
