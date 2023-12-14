package models

import (
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
}
