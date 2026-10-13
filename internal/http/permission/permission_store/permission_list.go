package permissionstore

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

func (s *PermissionStore) list(
	ctx context.Context,
	db *sql.DB,
	filter *permissiondto.PermissionFilter,
) ([]*permissiondto.Permission, int, error) {

	if filter.Page <= 0 {
		filter.Page = 1
	}

	if filter.PageSize <= 0 {
		filter.PageSize = 20
	}

	if filter.PageSize > 100 {
		filter.PageSize = 100
	}

	offset := (filter.Page - 1) * filter.PageSize

	args := []any{}
	where := []string{
		"is_deleted = FALSE",
	}

	if filter.Search != "" {
		args = append(args, "%"+filter.Search+"%")

		where = append(
			where,
			fmt.Sprintf(
				"CAST(id AS TEXT) ILIKE $%d",
				len(args),
			),
		)
	}

	whereClause := strings.Join(where, " AND ")

	countQuery := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM permission_master
		WHERE %s
	`, whereClause)

	var total int

	err := db.QueryRowContext(
		ctx,
		countQuery,
		args...,
	).Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	sortBy := "id"

	switch filter.SortBy {
	case "id":
		sortBy = "id"
	case "created_at":
		sortBy = "created_at"
	case "updated_at":
		sortBy = "updated_at"
	}

	sortOrder := "DESC"

	if strings.EqualFold(filter.SortOrder, "asc") {
		sortOrder = "ASC"
	}

	args = append(args, filter.PageSize)
	limitPosition := len(args)

	args = append(args, offset)
	offsetPosition := len(args)

	listQuery := fmt.Sprintf(`
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
		WHERE %s
		ORDER BY %s %s
		LIMIT $%d
		OFFSET $%d
	`, whereClause, sortBy, sortOrder, limitPosition, offsetPosition)

	rows, err := db.QueryContext(
		ctx,
		listQuery,
		args...,
	)

	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	permissions := make([]*permissiondto.Permission, 0)

	for rows.Next() {

		var permission permissiondto.Permission

		err := rows.Scan(
			&permission.ID,
			&permission.AllPerm,
			&permission.CreatePerm,
			&permission.ReadPerm,
			&permission.UpdatePerm,
			&permission.DeletePerm,
			&permission.CreatedBy,
			&permission.UpdatedBy,
			&permission.DeletedBy,
			&permission.CreatedAt,
			&permission.UpdatedAt,
			&permission.DeletedAt,
			&permission.IsDeleted,
		)

		if err != nil {
			return nil, 0, err
		}

		permissions = append(
			permissions,
			&permission,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return permissions, total, nil
}
