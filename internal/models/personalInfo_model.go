package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PersonalInfo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName   string             `bson:"first_name" json:"first_name"`
	LastName    string             `bson:"last_name" json:"last_name"`
	Email       string             `bson:"email" json:"email"`
	Phone       string             `bson:"phone" json:"phone"`
	Address     string             `bson:"address" json:"address"`
	Title       string             `bson:"title" json:"title"`
	LinkedInURL string             `bson:"linkedin_url" json:"linkedinUrl"`
	GitHubURL   string             `bson:"github_url" json:"githubUrl"`
	AboutMe     string             `bson:"about_me" json:"aboutMe"`
}
