package handler

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/rajeshbond/smart/internal/http/device/device_data/service"
)

type Handler struct {
	Service   *service.Service
	tokenAuth *jwtauth.JWTAuth
}

func NewService(service *service.Service, tokenAuth *jwtauth.JWTAuth) *Handler {
	return &Handler{
		Service:   service,
		tokenAuth: tokenAuth,
	}
}
