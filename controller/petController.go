package controller

import (
	"github.com/Strke12i/go_docker/models"
	"github.com/Strke12i/go_docker/service"
	"github.com/gin-gonic/gin"
)

type PetController struct {
	PetService service.PetService
}

func NewPetController(petService service.PetService) *PetController {
	return &PetController{PetService: petService}
}

func (controller *PetController) Create(context *gin.Context) {
	var pet models.Pet
	err := context.ShouldBindJSON(&pet)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = controller.PetService.Create(pet)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": pet})
}

func (controller *PetController) Update(context *gin.Context) {
	var pet models.Pet
	err := context.ShouldBindJSON(&pet)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = controller.PetService.Update(pet)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": pet})
}

func (controller *PetController) Delete(context *gin.Context) {
	var pet models.Pet
	err := context.ShouldBindJSON(&pet)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = controller.PetService.Delete(pet)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": pet})
}

func (controller *PetController) FindAll(context *gin.Context) {
	pets, err := controller.PetService.FindAll()
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": pets})
}

func (controller *PetController) FindById(context *gin.Context) {
	if id, ok := context.Params.Get("id"); ok {
		pet, err := controller.PetService.FindById(id)
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		context.JSON(200, gin.H{"data": pet})
	}
}
