package services

import (
	//"errors"
	"github.com/dasun/first_api/src/models"
	"github.com/dasun/first_api/src/repositories"
)

func GetAllUsers() []models.User {
	return repositories.GetAllUsers()
}

func GetUserByID(id int) (models.User, error) {
	return repositories.GetUserByID(id)
}

func CreateUser(newUser models.User) models.User {
	return repositories.CreateUser(newUser)
}

func DeleteUserByID(id int) error {
	return repositories.DeleteUserByID(id)
}

func UpdateUserByID(id int, updateUser models.User) error {
	return repositories.UpdateUserByID(id, updateUser)
}
