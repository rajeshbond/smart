package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (h *Handler) GetProductionLogByTenantIDAndDeviceID(
	w http.ResponseWriter,
	r *http.Request,
) {

	ctx := r.Context()

	//----------------------------------------------------------
	// Authentication
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

	//----------------------------------------------------------
	// Authorization
	//----------------------------------------------------------

	if !permission.CanCreateDevice(claims.Role) {

		response.Error(
			w,
			http.StatusForbidden,
			auth.PermissionDenied.Error(),
		)

		return
	}

	//----------------------------------------------------------
	// Decode Request
	//----------------------------------------------------------

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

	if req.Limit <= 0 {
		req.Limit = 1
	}

	//----------------------------------------------------------
	// Service
	//----------------------------------------------------------

	items, err := h.Service.GetProductionLogByTenantIDAndDeviceID(
		ctx,
		req,
	)

	if err != nil {

		response.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	if len(items) == 0 {

		response.Error(
			w,
			http.StatusNotFound,
			"production data not found",
		)

		return
	}

	//----------------------------------------------------------
	// Success
	//----------------------------------------------------------

	response.JSON(
		w,
		http.StatusOK,
		items,
	)
}
