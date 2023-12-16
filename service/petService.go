package service

import (
	"github.com/Strke12i/go_docker/models"
	"github.com/Strke12i/go_docker/repository"
	"github.com/go-playground/validator/v10"
)

type PetService interface {
	Create(pet models.Pet) error
	Update(pet models.Pet) error
	Delete(pet models.Pet) error
	FindAll() ([]models.Pet, error)
	FindById(id string) (models.Pet, error)
}

type PetServiceImpl struct {
	PetRepository repository.PetRepository
	Validate      *validator.Validate
}

func NewPetService(petRepository repository.PetRepository, validate *validator.Validate) *PetServiceImpl {
	return &PetServiceImpl{PetRepository: petRepository, Validate: validate}
}

func (service *PetServiceImpl) Create(pet models.Pet) error {
	err := service.Validate.Struct(pet)
	if err != nil {
		return err
	}
	return service.PetRepository.Save(pet)
}

func (service *PetServiceImpl) Update(pet models.Pet) error {
	err := service.Validate.Struct(pet)
	if err != nil {
		return err
	}
	return service.PetRepository.Update(pet)
}

func (service *PetServiceImpl) Delete(pet models.Pet) error {
	return service.PetRepository.Delete(pet)
}

func (service *PetServiceImpl) FindAll() ([]models.Pet, error) {
	result := service.PetRepository.FindAll()
	for _, pet := range result {
		err := service.Validate.Struct(pet)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (service *PetServiceImpl) FindById(id string) (models.Pet, error) {
	result, err := service.PetRepository.FindById(id)
	if err != nil {
		return result, err
	}

	err = service.Validate.Struct(result)
	if err != nil {
		return result, err
	}

	return result, nil
}
