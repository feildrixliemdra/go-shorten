package handler

import (
	"go-shorten/internal/appcontext"
	"go-shorten/internal/service"
)

type Handler struct {
	UserHandler IUserHandler
}

type Option struct {
	appcontext.Option
	Service *service.Service
}

func InitiateHandler(opt Option) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(opt),
	}
}
