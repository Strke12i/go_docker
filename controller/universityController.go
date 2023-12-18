package controller

import (
	"github.com/Strke12i/go_docker/models"
	"github.com/Strke12i/go_docker/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UniversityController struct {
	UniversityRepository *repository.UniversityRepository
}

func newUniversityController(DB *gorm.DB) *UniversityController {
	return &UniversityController{UniversityRepository: repository.NewUniversityRepository(DB)}
}

func (controller *UniversityController) GetAll(c *gin.Context) {
	universities, err := controller.UniversityRepository.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, universities)
}

func (controller *UniversityController) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required!"})
		return
	}

	university, err := controller.UniversityRepository.GetByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, university)
}

func (controller *UniversityController) Create(c *gin.Context) {
	var university models.University
	err := c.ShouldBindJSON(&university)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	university, err = controller.UniversityRepository.Create(university)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, university)
}

func (controller *UniversityController) Update(c *gin.Context) {
	var university models.University
	err := c.ShouldBindJSON(&university)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	university, err = controller.UniversityRepository.Update(university)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, university)
}

func (controller *UniversityController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required!"})
		return
	}

	err := controller.UniversityRepository.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "University deleted successfully!"})
}

func UniversityRegisterRoutes(route *gin.RouterGroup, DB *gorm.DB) {
	controller := newUniversityController(DB)
	route.GET("/universities", controller.GetAll)
	route.GET("/universities/:id", controller.GetByID)
	route.POST("/universities", controller.Create)
	route.PUT("/universities", controller.Update)
	route.DELETE("/universities/:id", controller.Delete)
}
