package controllers

import (
	"fmt"
	"net/http"
	"zunn/backend-api/database"
	"zunn/backend-api/helpers"
	"zunn/backend-api/models"
	"zunn/backend-api/structs"

	"github.com/gin-gonic/gin"
)

func validationErrors(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
		Success: false,
		Message: "Validation Errors",
		Errors: helpers.TranslateErrorMessage(err),
	})
}

func userNotFound(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, structs.ErrorResponse{
		Success: false,
		Message: "User Not Found",
		Errors: helpers.TranslateErrorMessage(err),
	})
}

func intervalServerError(c *gin.Context, message string, err error) {
	c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
		Success: false,
		Message: message,
		Errors: helpers.TranslateErrorMessage(err),
	})
}

func FindUser(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: 	true,
		Message: 	"List Data Users",
		Data:		users,
	})
}

func CreateUser(c *gin.Context) {
	var req = structs.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors(c, err)
		return
	}

	user := models.User{
		Name: req.Name,
		Username: req.Username,
		Email: req.Email,
		Password: helpers.HashPassword(req.Password),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		intervalServerError(c, "Failed To Create User", err)
		return
	}
	
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User Created Successfully",
		Data: structs.UserResponse{
			Id: 		user.ID,
			Name: 		user.Name,
			Username: 	user.Username,
			Email: 		user.Email,
			CreatedAt: 	user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: 	user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
	
}

func FindUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		userNotFound(c, err)
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User Found",
		Data: structs.UserResponse{
			Id: user.ID,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		userNotFound(c, err)
		return
	}

	var req = &structs.UserUpdateRequest{}

	if err := c.ShouldBindJSON(req); err != nil {
		validationErrors(c, err)
		return
	}

	fmt.Println("\nREQ: ", req)
	// return

	user.Name = req.Name
	user.Username = req.Username
	user.Email = req.Email

	if req.Password != "" {
		user.Password = helpers.HashPassword(req.Password)
	}

	// return

	if err := database.DB.Save(&user).Error; err != nil {
		intervalServerError(c, "Failed To Update User", err)
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User Updated Successfully",
		Data: structs.UserResponse{
			Id: user.ID,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func RemoveUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.Find(&user, id).Error; err != nil {
		userNotFound(c, err)
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		intervalServerError(c, "Failed To Remove User", err)
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User Removed Successfully",
	})
}