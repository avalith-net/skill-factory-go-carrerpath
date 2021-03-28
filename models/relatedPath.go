package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RelatadPath struct {
	RelationID primitive.ObjectID `bson:"_id,omitempty" json:"relationID,omitempty"`
	UserId     string             `bson:"userid" json:"userid,omitempty"`
	UserPathId string             `bson:"userpathid" json:"userpathid,omitempty"`
	Path       Path               `bson:"path" json:"path,omitempty"`
	// TechnicalSkills []TechnicalSkill `bson:"technicalSkills" json:"technicalSkills,omitempty"`
	// SoftSkills      []Skill          `bson:"softSkills" json:"softSkills,omitempty"`
}
