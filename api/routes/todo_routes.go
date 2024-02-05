package routes

import (
	"github.com/errorsolver/Todo-Golang/api/handlers"
	"github.com/errorsolver/Todo-Golang/api/middleware"
	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine) {
	todo := r.Group("/todo")

	todo.GET("/all", middleware.VerifJWT(), handlers.GetAllTodos)
	todo.GET("/", middleware.VerifJWT(), handlers.GetTodoByUserID)
	todo.POST("/", middleware.VerifJWT(), handlers.CreateTodo)
	todo.PUT("/", middleware.VerifJWT(), handlers.UpdateTodo)
	todo.DELETE("/:id", middleware.VerifJWT(), handlers.DeleteTodo)
}
