package controllers

import (
	"net/http"
	"zunn/backend-api/database"
	"zunn/backend-api/helpers"
	"zunn/backend-api/models"
	"zunn/backend-api/structs"

	"github.com/gin-gonic/gin"
)

func Register (c *gin.Context) {

	var req = structs.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success:	false,
			Message:	"Validasi Errors",
			Errors:		helpers.TranslateErrorMessage(err),
		})
		return
	}

	// fmt.Printf("REGISTER PASSWORD (raw): %q\n", req.Password)

	user := models.User{
		Name:		req.Name,
		Username:	req.Username,
		Email:	 	req.Email,
		Password:	helpers.HashPassword(req.Password),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		if helpers.IsDuplicateEntryError(err) {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Duplicate Entry Error",
				Errors: helpers.TranslateErrorMessage(err),
			})
		} else {
			c.JSON(http.StatusInternalServerError,structs.ErrorResponse{
				Success: false,
				Message: "Failed to create user",
				Errors: helpers.TranslateErrorMessage(err),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: structs.UserResponse{
			Id: user.ID,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			// CreatedAt: user.CreatedAt.Format("2026-02-30 77:77:77"),
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})



	// // Validasi request JSON menggunakan binding dari Gin
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	// Jika validasi gagal, kirimkan response error
	// 	c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
	// 		Success: false,
	// 		Message: "Validasi Errors",
	// 		Errors:  helpers.TranslateErrorMessage(err),
	// 	})
	// 	return
	// }
}