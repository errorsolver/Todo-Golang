package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Todo   string `json:"todo" gorm:"not null"`
	IsDone bool   `json:"isDone" gorm:"default:false"`
	UserID uint   `json:"userID" gorm:"not null"`
}
