package handler

import (
	"encoding/json"
<<<<<<< HEAD
=======
	"log"
>>>>>>> ac0246d3717753324e52eca15a89848498ea0cc7
	"net/http"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
<<<<<<< HEAD
	"github.com/rajeshbond/smart/internal/common/utils"
	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (h *Handler) SavePreviousHour(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	//----------------------------------------------------------
	// Authentication
	//----------------------------------------------------------

	claims, err := auth.MustUserClaims(ctx)
	if err != nil {
=======
	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

// ============================================================
// SAVE PREVIOUS HOUR
// ============================================================
//
// API:
//
// POST /proddata/assemblyhr
//
// The handler does:
//
// 1. Authentication
// 2. Authorization
// 3. Decode request
// 4. Call service
// 5. Return response
//
// Hourly calculation itself is done in the service/store.
//
// ============================================================

func (h *Handler) SavePreviousHour(
	w http.ResponseWriter,
	r *http.Request,
) {

	ctx := r.Context()

	// ========================================================
	// AUTHENTICATION
	// ========================================================

	claims, err := auth.MustUserClaims(ctx)

	if err != nil {

>>>>>>> ac0246d3717753324e52eca15a89848498ea0cc7
		response.Error(
			w,
			http.StatusUnauthorized,
			auth.UnAuthorised,
		)

		return
	}

<<<<<<< HEAD
	//----------------------------------------------------------
	// Authorization
	//----------------------------------------------------------
=======
	// ========================================================
	// AUTHORIZATION
	// ========================================================
>>>>>>> ac0246d3717753324e52eca15a89848498ea0cc7

	if !permission.ProductionLogViewwer(claims.Role) {

		response.Error(
			w,
			http.StatusForbidden,
			auth.PermissionDenied.Error(),
		)

		return
	}

<<<<<<< HEAD
	//----------------------------------------------------------
	// Decode Request
	//----------------------------------------------------------

	var req dto.SaveHourlyProductionRequest

	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
=======
	// ========================================================
	// DECODE REQUEST
	// ========================================================

	var req dto.SaveHourlyProductionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

>>>>>>> ac0246d3717753324e52eca15a89848498ea0cc7
		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
<<<<<<< HEAD
=======

>>>>>>> ac0246d3717753324e52eca15a89848498ea0cc7
		return
	}

	// ========================================================
<<<<<<< HEAD
	// TENANT ID FROM JWT
	// ========================================================
	//
	// Do NOT allow tenant_id from frontend/Postman.
	//
	// JWT is the source of tenant identity.
	tenantCode, err := utils.GetTenantCode(claims.EmployeeID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	req.TenantID = tenantCode
=======
	// BASIC REQUEST LOG
	// ========================================================

	log.Printf(
		"========== SAVE PREVIOUS HOUR REQUEST ==========\n"+
			"TenantID : %s\n"+
			"DeviceID : %s\n"+
			"MachineID: %s\n"+
			"Station  : %s\n"+
			"Variant  : %v\n"+
			"=================================================",
		req.TenantID,
		req.DeviceID,
		req.MachineID,
		req.Station,
		req.Variant,
	)
>>>>>>> ac0246d3717753324e52eca15a89848498ea0cc7

	// ========================================================
	// SERVICE
	// ========================================================

<<<<<<< HEAD
	item, err := h.Service.SavePreviousHour(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
=======
	item, err := h.Service.SavePreviousHour(
		ctx,
		req,
	)

	if err != nil {

		log.Printf(
			"❌ SavePreviousHour failed: %v",
			err,
		)

		response.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

>>>>>>> ac0246d3717753324e52eca15a89848498ea0cc7
	// ========================================================
	// SUCCESS
	// ========================================================

<<<<<<< HEAD
	response.JSON(w, http.StatusOK, item)

=======
	response.JSON(
		w,
		http.StatusOK,
		item,
	)
>>>>>>> ac0246d3717753324e52eca15a89848498ea0cc7
}
