package dao

import (
	"errors"

	"github.com/sbecker/gin-api-demo/models"
)

// fake database

var users []models.User

func InitDb() {
	users = append(users, models.User{DefaultModel: models.DefaultModel{ID: "abc", ObjectType: "user"}, AuthToken: "userA", Email: "userA@example.com", DOB: "01/23/1984", FavoriteCity: "Dallas"})
	users = append(users, models.User{DefaultModel: models.DefaultModel{ID: "xyz", ObjectType: "user"}, AuthToken: "userB", Email: "userB@example.com", DOB: "02/27/1991", FavoriteCity: "Portland", Admin: true})
}

func GetAllUsers(currentUser models.User) []models.User {
	// if admin, return all
	if currentUser.Admin {
		return users
	}

	// if not admin, only return itself
	var results []models.User
	for _, u := range users {
		if u.ID == currentUser.ID {
			results = append(results, u)
		}
	}
	return results
}

func GetUserByID(id string, currentUser models.User) (models.User, error) {
	// if not admin, make only itself findable
	if !currentUser.Admin {
		for _, u := range users {
			if u.ID == id && u.ID == currentUser.ID {
				return u, nil
			}
		}
		return models.User{}, errors.New("Not found")
	}

	for _, u := range users {
		if u.ID == id {
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
