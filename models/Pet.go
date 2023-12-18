package models

import (
	"gorm.io/gorm"
)

// Pet struct is a representation of a pet
// Contains the ID and the name of the pet
type Pet struct {
	gorm.Model
	ID           int        `json:"id" gorm:"primary_key"`
	Name         string     `json:"name" gorm:"not null"`
	UniversityID int        `json:"university_id" gorm:"not null"`
	University   University `json:"university" gorm:"foreignKey:UniversityID"`
}
