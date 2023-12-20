package models

// Pet struct is a representation of a pet
// Contains the ID and the name of the pet
type Pet struct {
	ID           int        `json:"id" gorm:"primary_key"`
	Name         string     `json:"name" gorm:"not null"`
	UniversityID int        `json:"university_id" gorm:"not null"`
	ProfessorCPF string     `json:"professor_cpf" gorm:"not null"`
	University   University `gorm:"foreignKey:UniversityID"`
	Professor    Professor  `gorm:"foreignKey:ProfessorCPF"`
}
