package machinestore

// ============================================================
// IMM SELECT COLUMNS
// ============================================================

const immSelectColumns = `
	sr,
	tenant_id,
	dept_id,
	machine_name,
	machine_no,
	machine_make,
	tie_bar_distance,
	platten_size,
	location_ring_size,
	created_by,
	updated_by,
	is_deleted,
	deleted_by,
	created_at,
	updated_at,
	deleted_at
`

// ============================================================
// CREATE
// ============================================================

const queryCreateIMM = `
	INSERT INTO imm (
		tenant_id,
		dept_id,
		machine_name,
		machine_no,
		machine_make,
		tie_bar_distance,
		platten_size,
		location_ring_size,
		created_by,
		updated_by
	)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10
	)
	RETURNING
	` + immSelectColumns

// ============================================================
// GET BY ID
// ============================================================

const queryGetIMMByID = `
	SELECT
	` + immSelectColumns + `
	FROM imm
	WHERE
		sr = $1
		AND tenant_id = $2
		AND is_deleted = FALSE
`

// ============================================================
// LIST
// ============================================================

const queryListIMM = `
	SELECT
	` + immSelectColumns + `
	FROM imm
	WHERE
		tenant_id = $1
		AND is_deleted = FALSE
`

// ============================================================
// COUNT
// ============================================================

const queryCountIMM = `
	SELECT COUNT(*)
	FROM imm
	WHERE
		tenant_id = $1
		AND is_deleted = FALSE
`

// ============================================================
// UPDATE
// ============================================================

const queryUpdateIMM = `
	UPDATE imm
	SET
		dept_id = $1,
		machine_name = $2,
		machine_no = $3,
		machine_make = $4,
		tie_bar_distance = $5,
		platten_size = $6,
		location_ring_size = $7,
		updated_by = $8,
		updated_at = CURRENT_TIMESTAMP
	WHERE
		sr = $9
		AND tenant_id = $10
		AND is_deleted = FALSE
	RETURNING
	` + immSelectColumns

// ============================================================
// DELETE
// ============================================================

const querySoftDeleteIMM = `
	UPDATE imm
	SET
		is_deleted = TRUE,
		deleted_by = $1,
		deleted_at = CURRENT_TIMESTAMP,
		updated_by = $1,
		updated_at = CURRENT_TIMESTAMP
	WHERE
		sr = $2
		AND tenant_id = $3
		AND is_deleted = FALSE
`

// ============================================================
// DEPARTMENT EXISTS
// ============================================================

const queryDepartmentExists = `
	SELECT EXISTS (
		SELECT 1
		FROM department_master
		WHERE
			id = $1
			AND tenant_id = $2
			AND is_deleted = FALSE
	)
`
