package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User is the user's model of the DB MongoDB
type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Role           string             `bson:"role" json:"role,omitempty"`
	Admin          bool               `bson:"admin" json:"admin,omitempty"`
	Name           string             `bson:"name" json:"name,omitempty"`
	Surname        string             `bson:"surname" json:"surname,omitempty"`
	DateOfBirth    time.Time          `bson:"dateOfBirth" json:"dateOfBirth,omitempty"`
	Email          string             `bson:"email" json:"email"`
	Password       string             `bson:"password" json:"password,omitempty"`
	ProfilePhoto   string             `bson:"profilePhoto" json:"profilePhoto,omitempty"`
	SeniorityLevel Seniority          `bson:"seniorityLevel" json:"seniorityLevel,omitempty"`
}
