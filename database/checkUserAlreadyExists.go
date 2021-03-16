package database

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	CheckUser checkUserInterface = &checkUser{}
)

type checkUser struct {
}

type checkUserInterface interface {
	CheckUserAlreadyExists(string) (models.User, bool, string)
}

//CheckUserAlreadyExists receives the email and check if it is in the db
func (user *checkUser) CheckUserAlreadyExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("careerpath")
	col := db.Collection("user")
	condition := bson.M{"email": email}
	var result models.User
	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
