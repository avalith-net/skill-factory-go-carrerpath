package models

//Seniority is the struct that defines the user's status
type Seniority struct {
	Name        string `bson:"name" json:"name,omitempty"`
	Level       string `bson:"level" json:"level,omitempty"`
	Description string `bson:"desciption" json:"description,omitempty"`
}
