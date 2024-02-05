package database

import (
	"os"

	"github.com/errorsolver/Todo-Golang/api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("PGSQL")
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("ERRORL: Failed to connect database")
	}

	db.AutoMigrate(
		&model.User{},
		&model.Todo{},
	)

	DB = db
}
