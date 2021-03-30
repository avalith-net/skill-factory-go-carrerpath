package models

type Notification struct {
	UserId  string `bson:"userid" json:"userid,omitempty"`
	Message string `bson:"message" json:"message,omitempty"`
}
