package controller

import (
	"net/http"

	"github.co/Sstvnnn/PRETPA-WEB/service"
	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (u *UserControllerImpl) CreateUser(context *gin.Context) {
	var request struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	
	// Parse the incoming JSON request body
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}
	
	message, err := u.userService.CreateUser(request.Name, request.Email, request.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating user",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": message,
	})
	
}