/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : update.go
 *
 * DESCRIPTION :
 * Update Device Handler
 *
 ******************************************************************************/

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
	"github.com/rajeshbond/smart/internal/http/device/device_master/mapper"
)

//------------------------------------------------------------------------------
// Update Device
//------------------------------------------------------------------------------

func (h *Handler) Update(
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

	if !permission.CanUpdateDevice(claims.Role) {

		response.Error(
			w,
			http.StatusForbidden,
			auth.PermissionDenied.Error(),
		)

		return
	}

	//----------------------------------------------------------------------
	// Parse ID
	//----------------------------------------------------------------------

	id, err := strconv.ParseInt(
		chi.URLParam(r, "id"),
		10,
		64,
	)

	if err != nil {

		response.Error(
			w,
			http.StatusBadRequest,
			"invalid device id",
		)

		return
	}

	//----------------------------------------------------------------------
	// Decode Request
	//----------------------------------------------------------------------

	var req dto.UpdateDeviceRequest

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
	// DTO -> Model
	//----------------------------------------------------------------------

	device := mapper.ToModelForUpdate(id, req, claims.UserID)

	//----------------------------------------------------------------------
	// Service
	//----------------------------------------------------------------------

	if err := h.service.Update(
		ctx,
		device,
	); err != nil {

		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	//----------------------------------------------------------------------
	// Response
	//----------------------------------------------------------------------

	response.JSON(w, http.StatusOK, mapper.ToUpdateResponse(device))
}
