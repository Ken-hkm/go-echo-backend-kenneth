package models

type PersonalInfo struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	LinkedInURL string `json:"linkedinUrl"`
	GitHubURL   string `json:"githubUrl"`
}
