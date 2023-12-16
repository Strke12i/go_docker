package controller

import (
	"github.com/Strke12i/go_docker/models"
	"github.com/Strke12i/go_docker/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PetController struct {
	PetRepository *repository.PetRepository
}

func newPetController(DB *gorm.DB) *PetController {
	return &PetController{PetRepository: repository.NewPetRepository(DB)}
}

func (controller *PetController) GetAll(c *gin.Context) {
	pets, err := controller.PetRepository.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, pets)
}

func (controller *PetController) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required!"})
		return
	}

	pet, err := controller.PetRepository.GetByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, pet)
}

func (controller *PetController) Create(c *gin.Context) {
	var pet models.Pet
	err := c.ShouldBindJSON(&pet)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pet, err = controller.PetRepository.Create(pet)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pet)
}

func (controller *PetController) Update(c *gin.Context) {
	var pet models.Pet
	err := c.ShouldBindJSON(&pet)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pet, err = controller.PetRepository.Update(pet)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pet)
}

func (controller *PetController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required!"})
		return
	}

	err := controller.PetRepository.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Pet deleted successfully!"})
}

func (controller *PetController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/pets", controller.GetAll)
	router.GET("/pets/:id", controller.GetByID)
	router.POST("/pets", controller.Create)
	router.PUT("/pets", controller.Update)
	router.DELETE("/pets/:id", controller.Delete)
}
