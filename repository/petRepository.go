package repository

import (
	"github.com/Strke12i/go_docker/models"
	"gorm.io/gorm"
)

type petRepository interface {
	GetAll() ([]models.Pet, error)
	GetByID(id string) (models.Pet, error)
	Create(pet models.Pet) (models.Pet, error)
	Update(pet models.Pet) (models.Pet, error)
	Delete(id string) error
}

type PetRepository struct {
	DB *gorm.DB
}

func NewPetRepository(DB *gorm.DB) *PetRepository {
	return &PetRepository{DB: DB}
}

func (repository *PetRepository) GetAll() ([]models.Pet, error) {
	var pets []models.Pet
	err := repository.DB.Find(&pets).Error
	if err != nil {
		return nil, err
	}
	return pets, nil
}

func (repository *PetRepository) GetByID(id string) (models.Pet, error) {
	var pet models.Pet
	err := repository.DB.Where("id = ?", id).First(&pet).Error
	if err != nil {
		return models.Pet{}, err
	}
	return pet, nil
}

func (repository *PetRepository) Create(pet models.Pet) (models.Pet, error) {
	err := repository.DB.Create(&pet).Error
	if err != nil {
		return models.Pet{}, err
	}
	return pet, nil
}

func (repository *PetRepository) Update(pet models.Pet) (models.Pet, error) {
	err := repository.DB.Save(&pet).Error
	if err != nil {
		return models.Pet{}, err
	}
	return pet, nil
}

func (repository *PetRepository) Delete(id string) error {
	var pet models.Pet
	err := repository.DB.Where("id = ?", id).Delete(&pet).Error
	if err != nil {
		return err
	}
	return nil
}
