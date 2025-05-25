package service

import (
	"github.co/Sstvnnn/PRETPA-WEB/model"
	"github.co/Sstvnnn/PRETPA-WEB/repository"
)

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserServiceImpl{
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (u *UserServiceImpl) CreateUser (name, email, password string) (string, error){
	user := &model.User{
		Name: name,
		Email: email,
		Password: password,
	}
	err := u.userRepo.CreateUser(user)
	if err != nil{
		return "error", err
	}
	return "success", nil
}