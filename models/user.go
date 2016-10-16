package models

// User is a record for a person who can login to the system
type User struct {
	DefaultModel
	Email        string `json:"email"`
	DOB          string `json:"dob"`
	FavoriteCity string `json:"favorite_city"`
	Admin        bool   `json:"admin"`
	AuthToken    string `json:"auth_token"`
}

// NewUser instatiates a new user
func NewUser() User {
	return User{}
}
