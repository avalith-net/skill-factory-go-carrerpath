package database

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type registerInterface interface {
	InsertRegister(user models.UserRegister) (string, bool, error)
}

type register struct{}

var (
	//RegisterService is the reference variable to the register interface
	RegisterService registerInterface
)

func init() {
	RegisterService = &register{}
}

//InsertRegister inserts the user data in the db
func (service *register) InsertRegister(user models.UserRegister) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("careerpath")
	col := db.Collection("user")
	user.Password, _ = EncryptPassword(user.Password)
	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
