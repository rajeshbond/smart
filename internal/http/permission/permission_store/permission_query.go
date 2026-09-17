package permissionstore

const (
	queryCreatePermission = `
		INSERT INTO permission_master (
			all_perm,
			create_perm,
			read_perm,
			update_perm,
			delete_perm,
			created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING
			id,
			all_perm,
			create_perm,
			read_perm,
			update_perm,
			delete_perm,
			created_by,
			updated_by,
			deleted_by,
			created_at,
			updated_at,
			deleted_at,
			is_deleted
	`

	queryGetPermissionByID = `
		SELECT
			id,
			all_perm,
			create_perm,
			read_perm,
			update_perm,
			delete_perm,
			created_by,
			updated_by,
			deleted_by,
			created_at,
			updated_at,
			deleted_at,
			is_deleted
		FROM permission_master
		WHERE id = $1
		  AND is_deleted = FALSE
	`

	queryUpdatePermission = `
		UPDATE permission_master
		SET
			all_perm = $1,
			create_perm = $2,
			read_perm = $3,
			update_perm = $4,
			delete_perm = $5,
			updated_by = $6,
			updated_at = NOW()
		WHERE id = $7
		  AND is_deleted = FALSE
		RETURNING
			id,
			all_perm,
			create_perm,
			read_perm,
			update_perm,
			delete_perm,
			created_by,
			updated_by,
			deleted_by,
			created_at,
			updated_at,
			deleted_at,
			is_deleted
	`

	queryDeletePermission = `
		UPDATE permission_master
		SET
			is_deleted = TRUE,
			deleted_by = $1,
			deleted_at = NOW()
		WHERE id = $2
		  AND is_deleted = FALSE
	`
)
