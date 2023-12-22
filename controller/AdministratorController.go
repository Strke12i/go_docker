package controller

import (
	"github.com/Strke12i/go_docker/models"
	"github.com/Strke12i/go_docker/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdministratorController struct {
	AdministratorRepository *repository.AdministratorRepository
}

// NewAdministratorController returns a new AdministratorController
func NewAdministratorController(DB *gorm.DB) *AdministratorController {
	return &AdministratorController{
		AdministratorRepository: repository.NewAdministratorRepository(DB),
	}
}

// GetAll returns all administrators
func (controller *AdministratorController) GetAll(c *gin.Context) {
	administrators, err := controller.AdministratorRepository.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, administrators)
}

// GetByID returns an administrator by id
func (controller *AdministratorController) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required!"})
		return
	}

	administrator, err := controller.AdministratorRepository.GetByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, administrator)
}

// Create creates a new administrator
func (controller *AdministratorController) Create(c *gin.Context) {
	var administrator models.Administrator
	err := c.ShouldBindJSON(&administrator)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = administrator.BeforeCreate()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	administrator, err = controller.AdministratorRepository.Create(administrator)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, administrator)
}

// Update updates an administrator
func (controller *AdministratorController) Update(c *gin.Context) {
	var administrator models.Administrator
	err := c.ShouldBindJSON(&administrator)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	administrator, err = controller.AdministratorRepository.Update(administrator)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, administrator)
}

// Delete deletes an administrator by id
func (controller *AdministratorController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required!"})
		return
	}

	err := controller.AdministratorRepository.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Administrator deleted successfully!"})
}
