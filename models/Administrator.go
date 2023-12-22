package models

import "golang.org/x/crypto/bcrypt"

type Administrator struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

func (a *Administrator) BeforeCreate() (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	a.Password = string(hashedPassword)
	return nil

}
