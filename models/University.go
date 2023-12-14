package models

import (
	"gorm.io/gorm"
)

// University struct is a representation of a university
type University struct {
	gorm.Model
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
}
