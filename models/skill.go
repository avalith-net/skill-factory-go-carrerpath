package models

//Skill is the struct for the softskill
type Skill struct {
	Name    string `bson:"name" json:"name,omitempty"`
	Level   string `bson:"level" json:"level,omitempty"`
	Blocked bool   `bson:"blocked" json:"blocked,omitempty"`
}
