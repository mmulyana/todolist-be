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
	router.POST("/api/users", middleware.AuthMiddleware(), controller.CreateUser)
	router.GET("/api/users/:id", middleware.AuthMiddleware(), controller.FindUser)
	router.PUT("/api/users/:id", middleware.AuthMiddleware(), controller.UpdateUser)
	router.DELETE("/api/users/:id", middleware.AuthMiddleware(), controller.DeleteUser)

	return router
}
