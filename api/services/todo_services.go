package services

import (
	"github.com/errorsolver/Todo-Golang/api/model"
	"github.com/errorsolver/Todo-Golang/api/repository"
)

func GetAllTodos() ([]model.Todo, error) {
	return repository.GetAllTodos()
}
func GetTodoByUserID(todo model.Todo) ([]model.Todo, error) {
	return repository.GetTodoByUserID(todo)
}
func CreateTodo(todo model.Todo) error {
	return repository.CreateTodo(todo)
}
func UpdateTodo(todo model.Todo) error {
	return repository.UpdateTodo(todo)
}
func DeleteTodo(todo model.Todo) error {
	return repository.DeleteTodo(todo)
}
