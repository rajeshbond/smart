package handler

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/rajeshbond/smart/internal/http/assembly_master/service"
)

type Handler struct {
	service   *service.Service
	tokenAuth *jwtauth.JWTAuth
}

func NewHandler(service *service.Service, tokenAuth *jwtauth.JWTAuth) *Handler {
	return &Handler{
		service:   service,
		tokenAuth: tokenAuth,
	}
}

// ============================================================
// CREATE
// POST /assembly-master
// ============================================================
