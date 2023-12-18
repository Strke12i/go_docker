package repository

import (
	"github.com/Strke12i/go_docker/models"
	"gorm.io/gorm"
)

type UniversityRepository struct {
	DB *gorm.DB
}

func NewUniversityRepository(DB *gorm.DB) *UniversityRepository {
	return &UniversityRepository{DB: DB}
}

func (repository *UniversityRepository) GetAll() ([]models.University, error) {
	var universities []models.University
	err := repository.DB.Find(&universities).Error
	if err != nil {
		return nil, err
	}
	return universities, nil
}

func (repository *UniversityRepository) GetByID(id string) (models.University, error) {
	var university models.University
	err := repository.DB.Where("id = ?", id).First(&university).Error
	if err != nil {
		return models.University{}, err
	}
	return university, nil
}

func (repository *UniversityRepository) Create(university models.University) (models.University, error) {
	err := repository.DB.Create(&university).Error
	if err != nil {
		return models.University{}, err
	}
	return university, nil
}

func (repository *UniversityRepository) Update(university models.University) (models.University, error) {
	err := repository.DB.Save(&university).Error
	if err != nil {
		return models.University{}, err
	}
	return university, nil
}

func (repository *UniversityRepository) Delete(id string) error {
	var university models.University
	err := repository.DB.Where("id = ?", id).Delete(&university).Error
	if err != nil {
		return err
	}
	return nil
}
