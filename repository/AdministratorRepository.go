package repository

import (
	"github.com/Strke12i/go_docker/models"
	"gorm.io/gorm"
)

type AdministratorRepository struct {
	DB *gorm.DB
}

func NewAdministratorRepository(DB *gorm.DB) *AdministratorRepository {
	return &AdministratorRepository{DB: DB}
}

func (repository *AdministratorRepository) GetAll() ([]models.Administrator, error) {
	var administrators []models.Administrator
	err := repository.DB.Find(&administrators).Error
	return administrators, err
}

func (repository *AdministratorRepository) GetByID(id string) (models.Administrator, error) {
	var administrator models.Administrator
	err := repository.DB.First(&administrator, id).Error
	return administrator, err
}

func (repository *AdministratorRepository) Create(administrator models.Administrator) (models.Administrator, error) {
	err := repository.DB.Create(&administrator).Error
	return administrator, err
}

func (repository *AdministratorRepository) Update(administrator models.Administrator) (models.Administrator, error) {
	err := repository.DB.Save(&administrator).Error
	return administrator, err
}

func (repository *AdministratorRepository) Delete(id string) error {
	err := repository.DB.Delete(&models.Administrator{}, id).Error
	return err
}
