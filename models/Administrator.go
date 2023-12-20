package models

type Administrator struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
