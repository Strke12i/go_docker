package routes

import (
	"github.com/Strke12i/go_docker/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(DB *gorm.DB) *gin.Engine {
	router := gin.Default()
	if gin.Mode() == gin.TestMode {
		router = gin.New()
	}

	//
	r := router.Group("/api")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	controller.PetRegisterRoutes(r, DB)

	return router

}
