package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (h *Handler) GetProductionLogByTenantIDAndDeviceID(w http.ResponseWriter, r *http.Request) {
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

	var req dto.GetProductionRequest

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
	// Fetch Fresh (latest) Production Log
	//----------------------------------------------------------------------

	resp, err := h.Service.GetProductionLogByTenantIDAndDeviceID(ctx, req.DeviceID, req.TenantID)

	if err != nil {

		if err == sql.ErrNoRows {
			response.Error(
				w,
				http.StatusNotFound,
				"production data not found",
			)
			return
		}

		response.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	response.JSON(
		w,
		http.StatusOK,
		resp,
	)

}
