package models

type RelatadPath struct {
	UserId          string           `bson:"userid" json:"userid"`
	UserPathId      string           `bson:"userpathid" json:"userpathid"`
	TechnicalSkills []TechnicalSkill `bson:"technicalSkills" json:"technicalSkills,omitempty"`
	SoftSkills      []Skill          `bson:"softSkills" json:"softSkills,omitempty"`
}
