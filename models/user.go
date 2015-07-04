package models

type User struct {
	DefaultModel
	Email        string `json:"email"`
	DOB          string `json:"dob"`
	FavoriteCity string `json:"favorite_city"`
	Admin        bool   `json:"admin"`
	AuthToken    string `json:"auth_token"`
}
