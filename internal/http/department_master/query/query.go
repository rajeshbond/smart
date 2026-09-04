package query

// ============================================================
// CREATE
// ============================================================

const CreateDepartment = `
	INSERT INTO department_master (
		department,
		created_by,
		updated_by
	)
	VALUES ($1, $2, $2)
	RETURNING
		id,
		department,
		created_by,
		updated_by,
		created_at,
		updated_at,
		is_deleted,
		deleted_at
`

// ============================================================
// GET ALL DEPARTMENTS
// ============================================================

const GetAllDepartments = `
	SELECT
		id,
		department,
		created_by,
		updated_by,
		created_at,
		updated_at,
		is_deleted,
		deleted_at
	FROM department_master
	WHERE is_deleted = FALSE
	ORDER BY department ASC
`

// ============================================================
// GET DEPARTMENT BY ID
// ============================================================

const GetDepartmentByID = `
	SELECT
		id,
		department,
		created_by,
		updated_by,
		created_at,
		updated_at,
		is_deleted,
		deleted_at
	FROM department_master
	WHERE id = $1
	  AND is_deleted = FALSE
`

// ============================================================
// UPDATE
// ============================================================

const UpdateDepartment = `
	UPDATE department_master
	SET
		department = $1,
		updated_by = $2,
		updated_at = NOW()
	WHERE id = $3
	  AND is_deleted = FALSE
	RETURNING
		id,
		department,
		created_by,
		updated_by,
		created_at,
		updated_at,
		is_deleted,
		deleted_at
`

// ============================================================
// SOFT DELETE
// ============================================================

const DeleteDepartment = `
	UPDATE department_master
	SET
		is_deleted = TRUE,
		deleted_at = NOW(),
		updated_at = NOW(),
		updated_by = $1
	WHERE id = $2
	  AND is_deleted = FALSE
	RETURNING id
`
