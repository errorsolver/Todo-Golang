package handlers

import (
	"net/http"
	"strconv"

	"github.com/errorsolver/Todo-Golang/api/model"
	"github.com/errorsolver/Todo-Golang/api/services"
	"github.com/gin-gonic/gin"
)

func GetAllTodos(c *gin.Context) {
	todos, err := services.GetAllTodos()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "GetAllTodos Success",
		"data":    todos})
}

func GetTodoByUserID(c *gin.Context) {
	var todo model.Todo

	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	todo.UserID = userID

	result, err := services.GetTodoByUserID(todo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "GetTodoByUserID Success",
		"data":    result})
}

func CreateTodo(c *gin.Context) {
	var todo model.Todo

	if err := c.BindJSON(&todo); err != nil {
		if err.Error() == "EOF" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Make sure you have entered data in the body",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	todo.UserID = userID

	if err := services.CreateTodo(todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "CreateTodo Success",
	})
}

func UpdateTodo(c *gin.Context) {
	var todo model.Todo

	if err := c.BindJSON(&todo); err != nil {
		if err.Error() == "EOF" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Make sure you have entered data in the body",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	todo.UserID = userID

	if err := services.UpdateTodo(todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "UpdateTodo Success",
	})
}

func DeleteTodo(c *gin.Context) {
	var todo model.Todo

	strID := c.Param("id")
	uintID, err := strconv.ParseUint(strID, 10, 0)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	todo.ID = uint(uintID)

	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	todo.UserID = userID

	if err := services.DeleteTodo(todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "DeleteTodo Success",
	})
}
