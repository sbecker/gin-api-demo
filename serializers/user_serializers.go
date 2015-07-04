package serializers

import "github.com/sbecker/gin-api-demo/models"

type UsersJSON struct {
	Users []models.User `json:"users"`
}

type UserJSON struct {
	User models.User `json:"user"`
}

type UsersSubsetJSON struct {
	Users []UserSubset `json:"users"`
}

type UserSubsetJSON struct {
	User UserSubset `json:"user"`
}

type UserSubset struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	DOB   string `json:"dob"`
}

func NewUserSubset(user models.User) UserSubset {
	return UserSubset{
		Id:    user.Id,
		Email: user.Email,
		DOB:   user.DOB,
	}
}

func NewUserSubsetJSON(user models.User) UserSubsetJSON {
	return UserSubsetJSON{
		User: UserSubset{
			Id:    user.Id,
			Email: user.Email,
			DOB:   user.DOB,
		},
	}
}

func NewUserJSON(user models.User) UserJSON {
	return UserJSON{User: user}
}

func NewUsersJSON(users []models.User) UsersJSON {
	return UsersJSON{Users: users}
}

func NewUsersSubsetJSON(users []models.User) UsersSubsetJSON {
	json := UsersSubsetJSON{Users: []UserSubset{}}
	for _, user := range users {
		json.Users = append(json.Users, NewUserSubset(user))
	}
	return json
}

func SerializeUsers(users []models.User, currentUser models.User) interface{} {
	if currentUser.Admin {
		return NewUsersJSON(users)
	} else {
		return NewUsersSubsetJSON(users)
	}
}

func SerializeUser(user models.User, currentUser models.User) interface{} {
	if currentUser.Admin {
		return NewUserJSON(user)
	} else {
		return NewUserSubsetJSON(user)
	}
}
