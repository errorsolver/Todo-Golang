package repository

import (
	"github.com/errorsolver/Todo-Golang/api/database"
	"github.com/errorsolver/Todo-Golang/api/model"
)

func GetAllTodos() ([]model.Todo, error) {
	var todos []model.Todo
	err := database.DB.Find(&todos).Error
	if err != nil {
		return []model.Todo{}, err
	}
	return todos, nil
}

func GetTodoByUserID(todo model.Todo) ([]model.Todo, error) {
	var todos []model.Todo
	err := database.DB.Find(&todos, "user_id = ?", todo.UserID).Error
	if err != nil {
		return []model.Todo{}, err
	}
	return todos, nil
}

func CreateTodo(todo model.Todo) error {
	err := database.DB.Create(&todo).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo model.Todo) error {
	if err := database.DB.Save(&todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTodo(todo model.Todo) error {
	if err := database.DB.Delete(&todo).Error; err != nil {
		return err
	}
	return nil
}
