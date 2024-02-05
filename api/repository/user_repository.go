package repository

import (
	"github.com/errorsolver/Todo-Golang/api/database"
	"github.com/errorsolver/Todo-Golang/api/model"
)

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := database.DB.Find(&users).Error
	if err != nil {
		return []model.User{}, err
	}
	return users, nil
}

func GetUserByID(id int) (model.User, error) {
	var user model.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
func GetUserByNamePass(loginData model.User) (model.User, error) {
	var user model.User
	result := database.DB.First(&user, "name = ? AND password = ?", loginData.Name, loginData.Password)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func CreateUser(user model.User) error {
	err := database.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user model.User) error {
	// if err := database.DB.Updates(&user).Error; err != nil {
	if err := database.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	if err := database.DB.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
