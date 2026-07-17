/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : delete.go
 *
 * DESCRIPTION :
 * Delete Device Service
 *
 ******************************************************************************/

package service

import (
	"context"
	"fmt"
)

//------------------------------------------------------------------------------
// Delete Device
//------------------------------------------------------------------------------

func (s *Service) Delete(
	ctx context.Context,
	id int64,
	updatedBy int64,
) error {

	//----------------------------------------------------------------------
	// Validation
	//----------------------------------------------------------------------

	if id <= 0 {
		return fmt.Errorf("invalid device id")
	}

	if updatedBy <= 0 {
		return fmt.Errorf("invalid updated by")
	}

	//----------------------------------------------------------------------
	// Delete Device
	//----------------------------------------------------------------------

	return s.store.Delete(
		ctx,
		id,
		updatedBy,
	)
}
