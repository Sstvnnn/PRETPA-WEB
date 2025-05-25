package service

type UserService interface {
	CreateUser(name, email, password string) (string, error)
}