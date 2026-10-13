package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	"github.com/rajeshbond/smart/internal/http/assembly_master/dto"
)

func (h *Handler) CreateAssembly(w http.ResponseWriter, r *http.Request) {
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
	if !permission.ProductionLogViewwer(claims.Role) {

		response.Error(
			w,
			http.StatusForbidden,
			auth.PermissionDenied.Error(),
		)

		return
	}

	var req dto.CreateAssemblyMasterRequest

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

	resp, err := h.service.CreateAssemblyMaster(ctx, req, claims)

	if err != nil {

		response.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	if resp == nil {

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
		resp,
	)

}
