package store

const createAssemblyQuery = `
		INSERT INTO public.assembly_master (
			tenant_id,
			machine_id,
			assembly_name,
			device_id,
			station,
			variant,
			hour_target_output,
			created_by,
			updated_by
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $8
		)
		RETURNING
			id,
			tenant_id,
			machine_id,
			assembly_name,
			device_id,
			station,
			variant,
			hour_target_output,
			is_active,
			is_deleted,
			created_by,
			updated_by,
			created_at,
			updated_at
	`

const getByIDQuery = `
		SELECT
			id,
			tenant_id,
			machine_id,
			assembly_name,
			device_id,
			station,
			variant,
			hour_target_output,
			is_active,
			is_deleted,
			created_by,
			updated_by,
			created_at,
			updated_at
		FROM public.assembly_master
		WHERE id = $1
		  AND tenant_id = $2
		  AND is_deleted = FALSE
	`

const listQuert = `
		SELECT
			id,
			tenant_id,
			machine_id,
			assembly_name,
			device_id,
			station,
			variant,
			hour_target_output,
			is_active,
			is_deleted,
			created_by,
			updated_by,
			created_at,
			updated_at
		FROM public.assembly_master
		WHERE tenant_id = $1
		  AND is_deleted = FALSE
		ORDER BY id ASC
	`

const updateQuery = `
UPDATE public.assembly_master
		SET
			machine_id = $1,
			assembly_name = $2,
			device_id = $3,
			station = $4,
			variant = $5,
			hour_target_output = $6,
			is_active = $7,
			updated_by = $8,
			updated_at = NOW()
		WHERE id = $9
		  AND tenant_id = $10
		  AND is_deleted = FALSE
		RETURNING
			id,
			tenant_id,
			machine_id,
			assembly_name,
			device_id,
			station,
			variant,
			hour_target_output,
			is_active,
			is_deleted,
			created_by,
			updated_by,
			created_at,
			updated_at
	`
