/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : create.go
 *
 * DESCRIPTION :
 * Create Device Handler
 *
 ******************************************************************************/

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
	"github.com/rajeshbond/smart/internal/http/device/device_master/mapper"
)

//------------------------------------------------------------------------------
// Create Device
//------------------------------------------------------------------------------

func (h *Handler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {

	ctx := r.Context()

	//----------------------------------------------------------------------
	// Authentication
	//----------------------------------------------------------------------

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

	//----------------------------------------------------------------------
	// Validate Request
	//----------------------------------------------------------------------

	if r.Body == nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"request body is required",
		)
		return
	}

	//----------------------------------------------------------------------
	// Decode Request
	//----------------------------------------------------------------------

	var req dto.CreateDeviceRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	//----------------------------------------------------------------------
	// Create Device
	//----------------------------------------------------------------------

	device, err := h.service.Create(
		ctx,
		req,
		claims.UserID,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	//----------------------------------------------------------------------
	// Success
	//----------------------------------------------------------------------

	response.JSON(
		w,
		http.StatusCreated,
		mapper.ToCreateResponse(device),
	)
}
