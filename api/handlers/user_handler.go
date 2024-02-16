package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/errorsolver/Todo-Golang/api/middleware"
	"github.com/errorsolver/Todo-Golang/api/model"
	"github.com/errorsolver/Todo-Golang/api/services"
	"github.com/errorsolver/Todo-Golang/api/utils"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
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
		"message": "Get all users success",
		"data":    users,
	})
}

func GetUserByID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	user, err := services.GetUserByID(userId)
	var errCode int
	if err != nil {
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
			"code":    errCode | http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"message": "Get user by ID Success",
		"data":    user,
	})
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
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

	var err error
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	if err := services.CreateUser(user); err != nil {
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
		"message": "Create user success",
		"data":    user,
	})
}

func UpdateUser(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
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
	user.ID = userID

	if err := services.UpdateUser(user); err != nil {
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
		"message": "Update user success",
		"data":    user,
	})
}

func DeleteUser(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	if err := services.DeleteUser(id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"message": fmt.Sprintf("Delete user ID = %d success", id),
		"data":    gin.H{},
	})
}

func Login(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
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
	hashPass, err := services.GetUserPass(user.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	if ok := utils.CheckPasswordHash(user.Password, hashPass); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   "user or password not correct",
		})
		return
	}

	token, err := services.GetUserByName(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	tokenCookie := utils.CreateJWTCookie(c, token)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"message": "Login success",
		"data":    tokenCookie,
	})
}

func Logout(c *gin.Context) {
	utils.DeleteJWTCookie(c)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"message": "Logout success",
		"data":    gin.H{},
	})
}

func getUserID(c *gin.Context) (uint, error) {
	result, err := middleware.GetJWTData(c)
	if err != nil {
		return 0, err
	}

	userID := uint(result["data"].(map[string]interface{})["id"].(float64))
	return userID, nil
}
