package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/errorsolver/Todo-Golang/api/model"
	"github.com/errorsolver/Todo-Golang/api/services"
	"github.com/gin-gonic/gin"
)

func GetAllTodos(c *gin.Context) {
	todos, err := services.GetAllTodos()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"message": "Get all todos success",
		"data":    todos,
	})
}

func GetTodoByUserID(c *gin.Context) {
	var todo model.Todo

	userID, err := getUserID(c)
	if err != nil {
		if err.Error() == "record not found" {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"success": false,
				"code":    http.StatusNotFound,
				"error":   err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	todo.UserID = userID

	result, err := services.GetTodoByUserID(todo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"message": "Get todo by user ID success",
		"data":    result,
	})
}

func CreateTodo(c *gin.Context) {
	var todo model.Todo

	if err := c.BindJSON(&todo); err != nil {
		if err.Error() == "EOF" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"success": false,
				"code":    http.StatusBadRequest,
				"error":   "Make sure you have entered data in the body",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	todo.UserID = userID

	if err := services.CreateTodo(todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"code":    http.StatusCreated,
		"message": "Create todo success",
		"data":    todo,
	})
}

func UpdateTodo(c *gin.Context) {
	var todo model.Todo

	if err := c.BindJSON(&todo); err != nil {
		if err.Error() == "EOF" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"success": false,
				"code":    http.StatusBadRequest,
				"error":   "Make sure you have entered data in the body",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	// todoID???
	if todo.ID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   "PLease input todo ID",
		})
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	todo.UserID = userID

	if err := services.UpdateTodo(todo); err != nil {
		if err.Error() == "record not found" {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"success": false,
				"code":    http.StatusNotFound,
				"error":   err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"message": "Update todo success",
		"data":    todo,
	})
}

func DeleteTodo(c *gin.Context) {
	var todo model.Todo

	strID := c.Param("id")
	uintID, err := strconv.ParseUint(strID, 10, 0)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	todo.ID = uint(uintID)

	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	todo.UserID = userID

	if err := services.DeleteTodo(todo); err != nil {
		if err.Error() == "record not found" {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"success": false,
				"code":    http.StatusNotFound,
				"error":   err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"message": fmt.Sprintf("Delete todo ID = %d success", userID),
		"data":    todo,
	})
}
