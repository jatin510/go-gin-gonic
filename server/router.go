package server

import (
	"gin-gonic-api/controllers"
	"gin-gonic-api/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.GetStatus)

	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("/v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.Retrieve)
		}
	}

	return router

}
