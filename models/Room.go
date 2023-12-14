package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Number       int        `json:"number" gorm:"primary_key"`
	Building     string     `json:"building" gorm:"primary_key"`
	UniversityID int        `json:"university_id" gorm:"not null"`
	WifiName     string     `json:"wifi_name" gorm:"not null"`
	University   University `gorm:"foreignKey:UniversityID"`
}
