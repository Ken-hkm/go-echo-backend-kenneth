package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Expertise struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Skills          []string           `bson:"skills" json:"skills"`
	TechnicalSkills []TechnicalSkill   `bson:"technical_skills" json:"technical_skills"`
}

type TechnicalSkill struct {
	SkillName string   `bson:"skill_name" json:"skill_name"`
	Items     []string `bson:"items" json:"items"`
}
