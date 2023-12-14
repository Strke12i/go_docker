package repository

import (
	"github.com/Strke12i/go_docker/models"
	"gorm.io/gorm"
)

type PetRepository interface {
	Save(pet models.Pet) error
	Update(pet models.Pet) error
	Delete(pet models.Pet) error
	FindAll() []models.Pet
	FindById(id uint) (models.Pet, error)
}

type PetRepositoryImpl struct {
	Db *gorm.DB
}

func NewPetRepository(db *gorm.DB) *PetRepositoryImpl {
	return &PetRepositoryImpl{Db: db}
}

func (repository *PetRepositoryImpl) Save(pet models.Pet) error {
	return repository.Db.Create(&pet).Error
}

func (repository *PetRepositoryImpl) Update(pet models.Pet) error {
	return repository.Db.Save(&pet).Error
}

func (repository *PetRepositoryImpl) Delete(pet models.Pet) error {
	return repository.Db.Delete(&pet).Error
}

func (repository *PetRepositoryImpl) FindAll() []models.Pet {
	var pets []models.Pet
	repository.Db.Find(&pets)
	return pets
}

func (repository *PetRepositoryImpl) FindById(id uint) (models.Pet, error) {
	var pet models.Pet
	err := repository.Db.Where("id = ?", id).Find(&pet).Error
	return pet, err
}
