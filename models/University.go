package models

// University struct is a representation of a university
type University struct {
	ID      int       `json:"id" gorm:"primary_key"`
	Name    string    `json:"name" gorm:"not null"`
	Student []Student `gorm:"foreignKey:UniversityID"`
}
