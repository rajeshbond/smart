package permissionhandler

import (
	"encoding/json"
	"net/http"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

func (h *PermissionHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {

	ctx := r.Context()

	claims, err := auth.MustUserClaims(ctx)

	if err != nil {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	if claims == nil {
		response.Error(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	//----------------------------------------------------------------------
	// Authorization
	//----------------------------------------------------------------------

	if !permission.IsXoomUser(claims.Role) {
		response.Error(
			w,
			http.StatusForbidden,
			auth.PermissionDenied.Error(),
		)
		return
	}

	// --------------------------------------------------
	// Decode request
	// --------------------------------------------------

	var req permissiondto.CreatePermissionRequest

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	// --------------------------------------------------
	// created_by comes from JWT
	// Do not trust client supplied created_by
	// --------------------------------------------------

	req.CreatedBy = &claims.UserID

	// --------------------------------------------------
	// Service
	// --------------------------------------------------

	permission, err := h.PermissionService.Create(
		ctx,
		&req,
	)

	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	// --------------------------------------------------
	// Response
	// --------------------------------------------------

	response.JSON(
		w,
		http.StatusCreated,
		permissiondto.PermissionResponse{
			Data: permission,
		},
	)
}
