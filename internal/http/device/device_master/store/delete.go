/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : delete.go
 *
 ******************************************************************************/

package store

import (
	"context"
	"database/sql"
)

func (s *Store) DeleteTx(
	ctx context.Context,
	tx *sql.Tx,
	id int64,
	updatedBy int64,
) error {

	result, err := tx.ExecContext(
		ctx,
		DeleteDevice,
		id,
		updatedBy,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
