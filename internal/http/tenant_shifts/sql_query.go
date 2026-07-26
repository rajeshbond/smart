package tenantshifts

const GetShiftName = `
SELECT
    shift_name
FROM tenant_shift
WHERE id = $1;
`
