/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : handler.go
 *
 * DESCRIPTION :
 * Device Master HTTP Handler
 *
 ******************************************************************************/

package handler

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/rajeshbond/smart/internal/http/device/device_master/service"
)

type Handler struct {
	service   *service.Service
	tokenAuth *jwtauth.JWTAuth
}

func NewHandler(
	service *service.Service,
	tokenAuth *jwtauth.JWTAuth,
) *Handler {

	return &Handler{
		service:   service,
		tokenAuth: tokenAuth,
	}
}
