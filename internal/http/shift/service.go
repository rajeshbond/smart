package shift

import (
	"context"
	"fmt"
	"strings"
)

type Service struct {
	Store *Store
}

func NewService(store *Store) *Service {
	return &Service{Store: store}
}

// ==========================================
// SINGLE CREATE (optional)
// ==========================================
func (s *Service) Create(ctx context.Context, userID int64, req CreateShiftTimingRequest) error {

	tx, err := s.Store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := s.Store.Create(ctx, tx, userID, req); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// ==========================================
// BULK CREATE (MAIN API)
// ==========================================
func (s *Service) CreateBulk(
	ctx context.Context,
	userID int64,
	req BulkCreateShiftRequest,
) error {

	if len(req) == 0 {
		return fmt.Errorf("empty request")
	}

	tx, err := s.Store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	tenantCache := make(map[string]int64)

	for _, shift := range req {

		//----------------------------------
		// Get Tenant
		//----------------------------------

		tenantID, ok := tenantCache[shift.TenantCode]
		if !ok {

			id, err := s.Store.GetTenantIDByCode(
				ctx,
				tx,
				shift.TenantCode,
			)

			if err != nil {
				return err
			}

			tenantID = id
			tenantCache[shift.TenantCode] = id
		}

		//----------------------------------
		// Parse New Shift
		//----------------------------------

		newShift, err := BuildShiftInterval(
			shift.ShiftName,
			shift.ShiftStart,
			shift.ShiftEnd,
		)

		if err != nil {
			return err
		}

		//----------------------------------
		// Existing Shifts
		//----------------------------------

		existing, err := s.Store.GetTenantShifts(
			ctx,
			tx,
			tenantID,
		)

		if err != nil {
			return err
		}

		//----------------------------------
		// Duplicate Name
		//----------------------------------

		for _, ex := range existing {

			if strings.EqualFold(
				ex.ShiftName,
				newShift.ShiftName,
			) {

				return fmt.Errorf(
					"shift '%s' already exists",
					newShift.ShiftName,
				)
			}
		}

		//----------------------------------
		// Overlap Check
		//----------------------------------

		if err := ValidateOverlap(
			existing,
			newShift,
		); err != nil {

			return err
		}

		//----------------------------------
		// Total Duration
		//----------------------------------

		if err := ValidateTotalDuration(
			existing,
			newShift,
		); err != nil {

			return err
		}

		//----------------------------------
		// Save
		//----------------------------------

		shiftID, err := s.Store.UpsertTenantShift(
			ctx,
			tx,
			tenantID,
			shift.ShiftName,
			userID,
		)

		if err != nil {
			return err
		}

		err = s.Store.InsertShiftTiming(
			ctx,
			tx,
			shiftID,
			shift.ShiftStart,
			shift.ShiftEnd,
			userID,
		)

		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
