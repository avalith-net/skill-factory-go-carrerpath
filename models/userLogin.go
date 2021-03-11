package models

import (
	"mime/multipart"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User is the user's model of the DB MongoDB
type User struct {
	ID             primitive.ObjectID    `bson:"_id,omitempty" form:"id"`
	Role           string                `bson:"role" form:"role,omitempty"`
	Leader         string                `bson:"leader" form:"leader,omitempty"`
	Admin          bool                  `bson:"admin" form:"admin,omitempty"`
	Name           string                `bson:"name" form:"name,omitempty"`
	Surname        string                `bson:"surname" form:"surname,omitempty"`
	DateOfBirth    time.Time             `bson:"dateOfBirth" form:"dateOfBirth,omitempty"`
	Email          string                `bson:"email" form:"email"`
	Password       string                `bson:"password" form:"password,omitempty"`
	ProfilePhoto   *multipart.FileHeader `bson:"profilePhoto" form:"profilePhoto,omitempty"`
	SeniorityLevel Seniority             `bson:"seniorityLevel" form:"seniorityLevel,omitempty"`
}
