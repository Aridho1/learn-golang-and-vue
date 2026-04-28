package controllers

import (
	"net/http"
	"zunn/backend-api/database"
	"zunn/backend-api/helpers"
	"zunn/backend-api/models"
	"zunn/backend-api/structs"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	var req = structs.UserLoginRequest{}
	var user = models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors: helpers.TranslateErrorMessage(err),
		})

		return
	}

	// fmt.Println("REQ:", req)

	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		// fmt.Println("USER NOT FOUND:", err)

		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "User Not Found",
			Errors: helpers.TranslateErrorMessage(err),
		})

		return
	}

	// fmt.Println("DB Password:", user.Password)
	// fmt.Println("INPUT Password:", req.Password)
	// fmt.Printf("INPUT Password (raw): %q\n", req.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid Password",
			Errors: helpers.TranslateErrorMessage(err),
		})

		return
	}

	token := helpers.GenerateToken(user.Username)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login Success",
		Data: structs.UserResponse{
			Id: user.ID,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
			Token: &token,
		},
	})
}