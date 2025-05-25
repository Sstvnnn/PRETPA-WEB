package main

import (
	"github.co/Sstvnnn/PRETPA-WEB/controller"
	"github.co/Sstvnnn/PRETPA-WEB/database"
	"github.co/Sstvnnn/PRETPA-WEB/repository"
	"github.co/Sstvnnn/PRETPA-WEB/router"
	"github.co/Sstvnnn/PRETPA-WEB/service"
)

func main() {
	// fmt.Println("TPA WEB 25-1")

	// var a string = "blabla"
	// fmt.Println(a)

	db := database.ConnectDatabase()
	userRepository := repository.NewUserRepository(db)

	// user := &model.User{
	// 	Name: "steven",
	// 	Email: "steven@gmail.com",
	// 	Password: "steven123",
	// }
	// err := userRepository.CreateUser(user)

	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	router.NewRouter(userController)
	// message, err := userService.CreateUser("keni", "keni@gmail.com", "keni123")
	// fmt.Println(message)

	// if err != nil{
	// 	panic(err)
	// }
}