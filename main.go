package main

import (
	"github.com/errorsolver/Todo-Golang/api/database"
	"github.com/errorsolver/Todo-Golang/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initializeApp() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	database.InitDB()
	return nil
}

func main() {
	err := initializeApp()
	if err != nil {
		panic("ERROR failed to initialize app: " + err.Error())
	}
	r := gin.Default()

	routes.UserRoutes(r)
	routes.TodoRoutes(r)

	r.Run(":8888")
}
