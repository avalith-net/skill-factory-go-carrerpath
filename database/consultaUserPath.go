package database

import (
	"context"
	"fmt"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultUserPath(path models.RelatadPath) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("careerpath")
	col := db.Collection("relation")

	condition := bson.M{
		"userid":     path.UserId,
		"userpathid": path.UserPathId,
	}
	//is necessary ?
	fmt.Println(path.UserId)
	fmt.Println(path.UserPathId)
	//

	var result models.RelatadPath
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}
	return true, nil
}
