package models

type RelatadPath struct {
	UserId string `bson:"userid" json:"userid"`
	UserPathId string `bson:"userpathid" json:"userpathid"`
	UserTechSkill TechnicalSkill `bson:"usertechskill" json:"usertechskill""`
	UserSoftSkill Skill `bson:"usersoftskill" json:"usersoftskill"`
}