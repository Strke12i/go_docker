package models

type Professor struct {
	CPF    string `json:"cpf" gorm:"primary_key;autoIncrement:false;size:11"`
	Name   string `json:"name" gorm:"not null"`
	Email  string `json:"email" gorm:"not null"`
	Passwd string `json:"password" gorm:"not null"`
}
