package repositories

import (
	"errors"

	"github.com/dasun/first_api/src/models"
)

var users = []models.User{
	{ID: 1, Username: "Dasun", Address: "Mirigama", Mobile: "0711234567", Age: 30, Email: "dasun.e.com"},
	{ID: 2, Username: "Amith", Address: "Colombo", Mobile: "0722222222", Age: 15, Email: "amith.e.com"},
	{ID: 3, Username: "Lasan", Address: "Gampaha", Mobile: "07188888888", Age: 25, Email: "lasan.e.com"},
}

func GetAllUsers() []models.User {
	return users
}

func GetUserByID(id int) (models.User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func CreateUser(newUser models.User) models.User {
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	return newUser
}

func DeleteUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func UpdateUserByID(id int, updateUser models.User) error {
	for i, u := range users {
		if u.ID == id {
			updateUser.ID = id
			users[i] = updateUser
			return nil
		}
	}
	return errors.New("user not found")
}
