package controllers

import (
	"net/http"
	"zunn/backend-api/database"
	"zunn/backend-api/models"
	"zunn/backend-api/structs"

	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: 	true,
		Message: 	"List Data Users",
		Data:		users,
	})
}