-- 000002_create_assembly_production_log.down.sql

------------------------------------------------------------------------------
-- Drop Indexes
------------------------------------------------------------------------------

DROP INDEX IF EXISTS idx_assembly_production_created_at;

DROP INDEX IF EXISTS idx_assembly_production_timestamp;

DROP INDEX IF EXISTS idx_assembly_production_station;

DROP INDEX IF EXISTS idx_assembly_production_machine_id;

DROP INDEX IF EXISTS idx_assembly_production_device_id;

DROP INDEX IF EXISTS idx_assembly_production_customer_id;

DROP INDEX IF EXISTS idx_assembly_production_tenant_id;

------------------------------------------------------------------------------
-- Drop Table
------------------------------------------------------------------------------

DROP TABLE IF EXISTS assembly_production_log;