package dao

import (
	"errors"

	"github.com/sbecker/gin-api-demo/models"
)

// fake database

var users []models.User

func InitDb() {
	users = append(users, models.User{DefaultModel: models.DefaultModel{Id: "abc", ObjectType: "user"}, AuthToken: "userA", Email: "userA@example.com", DOB: "01/23/1984", FavoriteCity: "Dallas"})
	users = append(users, models.User{DefaultModel: models.DefaultModel{Id: "xyz", ObjectType: "user"}, AuthToken: "userB", Email: "userB@example.com", DOB: "02/27/1991", FavoriteCity: "Portland", Admin: true})
}

func GetAllUsers(currentUser models.User) []models.User {
	// if admin, return all
	if currentUser.Admin {
		return users
	}

	// if not admin, only return itself
	var results []models.User
	for _, u := range users {
		if u.Id == currentUser.Id {
			results = append(results, u)
		}
	}
	return results
}

func GetUserById(id string, currentUser models.User) (models.User, error) {
	// if not admin, make only itself findable
	if !currentUser.Admin {
		for _, u := range users {
			if u.Id == id && u.Id == currentUser.Id {
				return u, nil
			}
		}
		return models.User{}, errors.New("Not found")
	}

	for _, u := range users {
		if u.Id == id {
			return u, nil
		}
	}
	return models.User{}, errors.New("Not found")
}

func GetUserByAuthToken(token string) (models.User, error) {
	for _, u := range users {
		if u.AuthToken == token {
			return u, nil
		}
	}
	return models.User{}, errors.New("Not found")
}
