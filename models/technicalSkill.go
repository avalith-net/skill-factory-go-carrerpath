package models

//TechnicalSkill is the struct of the main technial skill which it has an slice of tools
type TechnicalSkill struct {
	Name    string      `bson:"name" json:"name,omitempty"`
	Blocked bool        `bson:"blocked" json:"blocked,omitempty"`
	Tools   []SkillTool `bson:"tools" json:"tools,omitempty"`
}
