package router

import (
	"github.co/Sstvnnn/PRETPA-WEB/controller"
	"github.com/gin-gonic/gin"
)


func NewRouter(userController controller.UserController) {
	router := gin.Default()
	router.POST("/create-user", userController.CreateUser)
	router.Run()
}