package routes

import (
	"mmulyana/todolist-be/controller"
	"mmulyana/todolist-be/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	router.POST("/api/register", controller.Register)
	router.POST("/api/login", controller.Login)

	router.GET("/api/users", middleware.AuthMiddleware(), controller.FindUsers)

	return router
}
