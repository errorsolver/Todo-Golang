package handlers

import (
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "GetAllUsers Success",
		"data":    users})
}

func GetUserByID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := services.GetUserByID(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "GetUserByID Success",
		"data":    user})
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
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

	if err := services.CreateUser(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "CreateUser Success",
	})
}

func UpdateUser(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
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
	user.ID = userID

	if err := services.UpdateUser(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "UpdateUser Success",
	})
}

func DeleteUser(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := services.DeleteUser(id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "DeleteUser Success",
	})
}

func Login(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
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

	token, err := services.GetUserByNamePass(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	utils.CreateJWTCookie(c, token)

	c.JSON(http.StatusOK, gin.H{"message": "Login success"})
}

func Logout(c *gin.Context) {
	utils.DeleteJWTCookie(c)
}

func getUserID(c *gin.Context) (uint, error) {
	result, err := middleware.GetJWTData(c)
	if err != nil {
		return 0, err
	}

	userID := uint(result["data"].(map[string]interface{})["id"].(float64))
	return userID, nil
}
