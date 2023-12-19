package models

type Student struct {
	CPF          string     `json:"cpf" gorm:"primary_key;autoIncrement:false;size:11"`
	PetID        int        `json:"pet_id"`
	UniversityID int        `json:"university_id"`
	Name         string     `json:"name" gorm:"not null"`
	Email        string     `json:"email" gorm:"not null"`
	Passwd       string     `json:"password" gorm:"not null"`
	Pet          Pet        `gorm:"foreignKey:PetID"`
	University   University `gorm:"foreignKey:UniversityID"`
}
