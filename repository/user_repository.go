package repository

import "github.co/Sstvnnn/PRETPA-WEB/model"

type UserRepository interface {
	CreateUser(user *model.User) error
}