package service

import (
	"go-shorten/internal/appcontext"
	"go-shorten/internal/repository"
)

type Service struct {
	UserService IUserService
}

type Option struct {
	appcontext.Option
	Repository *repository.Repository
}

func InitiateService(option Option) *Service {
	return &Service{
		UserService: NewUserService(option),
	}
}
