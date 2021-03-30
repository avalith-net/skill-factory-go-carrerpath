package models

type Notification struct {
	UserId    string   `bson:"userid" json:"userid,omitempty"`
	Name      string   `bson:"name" json:"name,omitempty"`
	TechSkill []string `bson:"tech_skill" json:"tech_skill,omitempty"`
	SoftSkill []string `bson:"soft_skill" json:"soft_skill,omitempty"`
	Message   string   `bson:"message" json:"message,omitempty"`
}
