package models

//SkillTool is the struct of any skill and subskill
type SkillTool struct {
	Name    string  `bson:"name" json:"name,omitempty"`
	Blocked bool    `bson:"blocked" json:"blocked,omitempty"`
	Tools   []Skill `bson:"tool" json:"tool"`
}
