/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : delete.go
 *
 * DESCRIPTION :
 * Delete Device
 *
 ******************************************************************************/

package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
)

func (h *Handler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {

	ctx := r.Context()

	//----------------------------------------------------------------------
	// Authentication
	//----------------------------------------------------------------------

	claims, err := auth.MustUserClaims(ctx)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, auth.UnAuthorised)
		return
	}

	//----------------------------------------------------------------------
	// Authorization
	//----------------------------------------------------------------------

	if !permission.CanDeleteDevice(claims.Role) {
		response.Error(w, http.StatusForbidden, auth.PermissionDenied.Error())
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
		response.Error(w, http.StatusBadRequest, "invalid device id")
		return
	}

	//----------------------------------------------------------------------
	// Service
	//----------------------------------------------------------------------

	err = h.service.Delete(
		ctx,
		id,
		claims.UserID,
	)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	//----------------------------------------------------------------------
	// Response
	//----------------------------------------------------------------------

	response.JSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": "Device deleted successfully",
		},
	)
}
