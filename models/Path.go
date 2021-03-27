package models

//Path is the struct of the career path, where we have a description of the type of path and 2 slices
// each of them will have different technologies according to the path's description.
type Path struct {
	Description     string           `bson:"description" json:"description,omitempty"` //here is the type of path
	TechnicalSkills []TechnicalSkill `bson:"technicalSkills" json:"technicalSkills,omitempty"`
	SoftSkills      []Skill          `bson:"softSkills" json:"softSkills,omitempty"`
}
