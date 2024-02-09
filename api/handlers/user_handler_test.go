package handlers

import (
	"net/http"
	"testing"

	"github.com/errorsolver/Todo-Golang/api/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repoMock struct {
	mock.Mock
}

func (r *repoMock) GetAllTodos() ([]model.Todo, error) {
	return []model.Todo{}, nil
}

func GetAllTodosSuccess(t *testing.T) {
	data:=map[string]interface{}{
		"todo": "Do something",
		"isDOne": false,
		"userID": 1,
	}
	GetAllTodosServices := repoMock{}
	GetAllTodosServices.On("").Return(
		data, "error",
	)

	// result := GetAllTodos(c)

	expected := gin.H{
		"success": true,
		"code":    http.StatusOK,
		"message": "Get all todos success",
		"data":    gin.H{},
	}
	actual := ""

	assert.Equal(t, expected, actual)
}
