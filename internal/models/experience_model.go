package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Experience struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Company     string             `bson:"company" json:"company"`
	Location    string             `bson:"location" json:"location"`
	StartDate   string             `bson:"start_date" json:"start_date"`
	EndDate     string             `bson:"end_date" json:"end_date"`
	Description []ExperienceDetail `bson:"description" json:"description"`
}

type ExperienceDetail struct {
	Role    string   `bson:"role" json:"role"`
	Details []string `bson:"details" json:"details"`
}
