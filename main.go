package main

import (
	"github.com/errorsolver/Todo-Golang/api/database"
	"github.com/errorsolver/Todo-Golang/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("ERROR fail load env")
	}
	database.InitDB()

	r := gin.Default()

	routes.UserRoutes(r)
	routes.TodoRoutes(r)

	r.Run(":8888")
}
