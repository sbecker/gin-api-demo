package serializers

import "github.com/sbecker/gin-api-demo/models"

type UsersJSON struct {
	ListMeta
	Values []models.User `json:"data"`
}

type UsersSubsetJSON struct {
	ListMeta
	Values []UserSubset `json:"data"`
}

type UserSubset struct {
	models.DefaultModel
	Email string `json:"email"`
	DOB   string `json:"dob"`
}

func NewUserSubset(user models.User) UserSubset {
	return UserSubset{
		DefaultModel: user.DefaultModel,
		Email:        user.Email,
		DOB:          user.DOB,
	}
}

func NewUsersJSON(users []models.User, URL string) UsersJSON {
	return UsersJSON{
		Values: users,
		ListMeta: ListMeta{
			ObjectType: "list",
			URL:        URL,
			Count:      uint16(len(users)),
		},
	}
}

func NewUsersSubsetJSON(users []models.User, URL string) UsersSubsetJSON {
	json := UsersSubsetJSON{
		Values: []UserSubset{},
		ListMeta: ListMeta{
			ObjectType: "list",
			URL:        URL,
			Count:      uint16(len(users)),
		},
	}
	for _, user := range users {
		json.Values = append(json.Values, NewUserSubset(user))
	}
	return json
}

func SerializeUsers(users []models.User, currentUser models.User, URL string) interface{} {
	if currentUser.Admin {
		return NewUsersJSON(users, URL)
	}
	return NewUsersSubsetJSON(users, URL)
}

func SerializeUser(user models.User, currentUser models.User) interface{} {
	if currentUser.Admin {
		return user
	}
	return NewUserSubset(user)
}
