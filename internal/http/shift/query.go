package shift

const GetShiftTimingByTenant = `
SELECT
    ts.id,
    ts.shift_name,
    st.shift_start,
    st.shift_end
FROM tenant_shift ts
JOIN shift_timing st
    ON st.tenant_shift_id = ts.id
WHERE ts.tenant_id = $1
ORDER BY st.shift_start;
`
