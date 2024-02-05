package model

import (
	"gorm.io/gorm"
)

type User struct {
	// tambah anti duplikat
	gorm.Model
	Name     string `gorm:"unique;not null" json:"name"`
	Password string `json:"pass" gorm:"not null"`
	Todo     []Todo
}
