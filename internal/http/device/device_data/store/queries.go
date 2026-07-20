package store

const GetLatestProductionByDevice = `
SELECT
    tenant_id,
    customer_id,
    device_id,
    machine_id,
    station,
    production_count,
    cycle_time_sec,
    production_time,
    created_at
FROM assembly_production_log
WHERE tenant_id = $1
  AND device_id = $2
ORDER BY created_at DESC
LIMIT 5;
`
