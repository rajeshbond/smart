/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : register_mqtt.go
 *
 ******************************************************************************/

package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
)

func (h *Handler) RegisterMQTTUsername(
	w http.ResponseWriter,
	r *http.Request,
) {

	ctx := r.Context()

	//----------------------------------------------------------
	// Request
	//----------------------------------------------------------

	req := dto.RegisterMQTTRequest{

		DeviceID: chi.URLParam(
			r,
			"deviceID",
		),
	}

	//----------------------------------------------------------
	// Logged-in User
	//----------------------------------------------------------

	claims, err := auth.MustUserClaims(ctx)
	if err != nil {
		response.Error(
			w,
			http.StatusUnauthorized,
			auth.UnAuthorised,
		)
		return
	}

	//----------------------------------------------------------------------
	// Authorization
	//----------------------------------------------------------------------

	if !permission.CanCreateDevice(claims.Role) {
		response.Error(
			w,
			http.StatusForbidden,
			auth.PermissionDenied.Error(),
		)
		return
	}

	//----------------------------------------------------------
	// Service
	//----------------------------------------------------------

	resp, err := h.service.RegisterMqttUsername(

		ctx,

		req,

		claims.UserID,
	)

	if err != nil {

		response.Error(w, http.StatusBadRequest, err.Error())

		return
	}

	//----------------------------------------------------------
	// Success
	//----------------------------------------------------------

	response.JSON(

		w,

		http.StatusOK,

		resp,
	)
}
