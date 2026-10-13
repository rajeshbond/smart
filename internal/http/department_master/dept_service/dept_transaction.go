package deptservice

import (
	"context"
	"database/sql"
)

// ============================================================
// WITH TRANSACTION
// ============================================================

func (s *DeptService) withTransaction(
	ctx context.Context,
	fn func(tx *sql.Tx) error,
) error {

	tx, err := s.DeptStore.BeginTx(ctx)

	if err != nil {
		return err
	}

	// ============================================================
	// AUTOMATIC ROLLBACK
	// ============================================================

	defer func() {
		_ = tx.Rollback()
	}()

	// ============================================================
	// OPERATIONS
	// ============================================================

	if err := fn(tx); err != nil {
		return err
	}

	// ============================================================
	// COMMIT
	// ============================================================

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
