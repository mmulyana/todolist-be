package controller

import (
	"mmulyana/todolist-be/database"
	"mmulyana/todolist-be/models"
	"mmulyana/todolist-be/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Lists Data Users",
		Data:    users,
	})
}
