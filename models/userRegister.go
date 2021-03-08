package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserRegister is the database user register model
type UserRegister struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name           string             `bson:"name" json:"name,omitempty"`
	Surname        string             `bson:"surname" json:"surname,omitempty"`
	Email          string             `bson:"email" json:"email,omitempty" binding:"required" `
	Password       string             `bson:"password" json:"password,omitempty" binding:"required"`
	ProfilePicture string             `bson:"profile_picture" json:"profile_picture,omitempty"`
}
