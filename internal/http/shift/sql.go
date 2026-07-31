package shift

const GetShiftByTenant = `
SELECT
    ts.id,
    ts.shift_name,
    TO_CHAR(st.shift_start,'HH24:MI') AS shift_start,
    TO_CHAR(st.shift_end,'HH24:MI') AS shift_end
FROM tenant_shift ts
INNER JOIN shift_timing st
    ON ts.id = st.tenant_shift_id
WHERE ts.tenant_id = $1
ORDER BY st.shift_start;
`

// const GetShiftByTenant = `
// SELECT
//     ts.id,
//     ts.shift_name,
//     TO_CHAR(st.shift_start, 'HH24:MI') AS shift_start,
//     TO_CHAR(st.shift_end, 'HH24:MI') AS shift_end
// FROM tenant_shift ts
// INNER JOIN shift_timing st
//     ON st.tenant_shift_id = ts.id
// WHERE ts.tenant_id = $1
// ORDER BY
//     st.shift_start;

// const GetShiftByTenant = `
// SELECT
//     ts.id,
//     ts.shift_name,
//     st.shift_start,
//     st.shift_end,
//     st.weekday
// FROM tenant_shift ts
// INNER JOIN shift_timing st
//     ON st.tenant_shift_id = ts.id
// WHERE ts.tenant_id = $1
// ORDER BY
//     st.weekday,
//     st.shift_start;
// `
