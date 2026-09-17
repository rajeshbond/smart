package permissionstore

import (
	"context"
	"database/sql"
)

func (s *PermissionStore) delete(
	ctx context.Context,
	tx *sql.Tx,
	id int64,
	deletedBy *int64,
) error {

	result, err := tx.ExecContext(
		ctx,
		queryDeletePermission,
		deletedBy,
		id,
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
