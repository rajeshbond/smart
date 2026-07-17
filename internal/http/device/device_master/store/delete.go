/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : delete.go
 *
 * DESCRIPTION :
 * Soft Delete Device
 *
 ******************************************************************************/

package store

import (
	"context"
	"database/sql"
	"fmt"
)

// -----------------------------------------------------------------------------
// Delete Device
// -----------------------------------------------------------------------------
//
// Soft Deletes the device.
//
// UPDATE device_master
// SET
//      is_deleted = TRUE,
//      updated_by = ?,
//      updated_at = CURRENT_TIMESTAMP
//
// WHERE
//      id = ?
//      AND is_deleted = FALSE;
//
// -----------------------------------------------------------------------------

func (s *Store) Delete(
	ctx context.Context,
	id int64,
	updatedBy int64,
) error {

	result, err := s.db.ExecContext(
		ctx,
		DeleteDevice,
		id,
		updatedBy,
	)

	if err != nil {
		return fmt.Errorf("delete device: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
