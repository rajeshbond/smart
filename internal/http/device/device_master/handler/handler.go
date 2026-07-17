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
	"github.com/rajeshbond/smart/internal/http/device/device_master/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(
	service *service.Service,
) *Handler {

	return &Handler{
		service: service,
	}
}
