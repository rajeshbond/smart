/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : list.go
 *
 * DESCRIPTION :
 * List Devices
 *
 ******************************************************************************/

package handler

import (
	"net/http"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
)

//------------------------------------------------------------------------------
// List Devices
//------------------------------------------------------------------------------

func (h *Handler) List(
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

	if !permission.HasPermission(
		claims.Role,
		permission.DeviceList,
	) {
		response.Error(
			w,
			http.StatusForbidden,
			auth.PermissionDenied.Error(),
		)
		return
	}

	//----------------------------------------------------------------------
	// Service
	//----------------------------------------------------------------------

	items, err := h.service.List(ctx)
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
		http.StatusOK,
		items,
	)
}
