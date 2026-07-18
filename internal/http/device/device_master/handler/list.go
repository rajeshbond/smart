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
	"strconv"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	"github.com/rajeshbond/smart/internal/http/device/device_master/dto"
	"github.com/rajeshbond/smart/internal/http/device/device_master/mapper"
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
	// Build Filter
	//----------------------------------------------------------------------

	filter := dto.DeviceFilter{
		Search:            r.URL.Query().Get("search"),
		DeviceStatus:      r.URL.Query().Get("device_status"),
		CommunicationType: r.URL.Query().Get("communication_type"),
	}

	// Page
	if page := r.URL.Query().Get("page"); page != "" {
		if p, err := strconv.Atoi(page); err == nil {
			filter.Page = p
		}
	}

	// Page Size
	if pageSize := r.URL.Query().Get("page_size"); pageSize != "" {
		if ps, err := strconv.Atoi(pageSize); err == nil {
			filter.PageSize = ps
		}
	}

	// Active
	if active := r.URL.Query().Get("is_active"); active != "" {
		if v, err := strconv.ParseBool(active); err == nil {
			filter.IsActive = &v
		}
	}

	//----------------------------------------------------------------------
	// Defaults
	//----------------------------------------------------------------------

	if filter.Page <= 0 {
		filter.Page = 1
	}

	if filter.PageSize <= 0 {
		filter.PageSize = 20
	}

	//----------------------------------------------------------------------
	// Service
	//----------------------------------------------------------------------

	devices, err := h.service.List(
		ctx,
		filter,
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
	// Response
	//----------------------------------------------------------------------

	response.JSON(
		w,
		http.StatusOK,
		mapper.ToListItems(devices),
	)
}
