package routes

import (
	"github.com/errorsolver/Todo-Golang/api/handlers"
	"github.com/errorsolver/Todo-Golang/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/login", handlers.Login)
	r.GET("/logout", handlers.Logout)

	user := r.Group("/user")

	user.POST("/register", handlers.CreateUser)
	user.GET("/all", middleware.VerifJWT(), handlers.GetAllUsers)
	user.GET("/:id", middleware.VerifJWT(), handlers.GetUserByID)
	user.PUT("/", middleware.VerifJWT(), handlers.UpdateUser)
	user.DELETE("/", middleware.VerifJWT(), handlers.DeleteUser)
}
