package services

import (
	"github.com/errorsolver/Todo-Golang/api/middleware"
	"github.com/errorsolver/Todo-Golang/api/model"
	"github.com/errorsolver/Todo-Golang/api/repository"
)

func GetUserByNamePass(user model.User) (string, error) {
	result, err := repository.GetUserByNamePass(user)
	if err != nil {
		return "", err
	}

	name, id := result.Name, result.ID
	token, err := middleware.CreateJwt(name, id)
	return token, err
}

func GetAllUsers() ([]model.User, error) {
	return repository.GetAllUsers()
}
func GetUserByID(id int) (model.User, error) {
	return repository.GetUserByID(id)
}

func CreateUser(user model.User) error {
	return repository.CreateUser(user)
}
func UpdateUser(user model.User) error {
	return repository.UpdateUser(user)
}
func DeleteUser(id uint) error {
	return repository.DeleteUser(id)
}
