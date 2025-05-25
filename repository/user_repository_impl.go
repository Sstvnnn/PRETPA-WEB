package repository

import (
	"github.co/Sstvnnn/PRETPA-WEB/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl{
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) CreateUser(user *model.User) error{
	err := u.db.Create(user).Error
	return err
}