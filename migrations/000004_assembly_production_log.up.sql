CREATE TABLE IF NOT EXISTS assembly_production_log (
    id BIGSERIAL PRIMARY KEY,
    tenant_id VARCHAR(50) NOT NULL,
    customer_id VARCHAR(50),
    device_id VARCHAR(50) NOT NULL,
    machine_id VARCHAR(50) NOT NULL,
    station VARCHAR(20) NOT NULL,
    production_count BIGINT NOT NULL,
    cycle_time_sec NUMERIC(10, 3),
    production_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_apl_tenant ON assembly_production_log (tenant_id);

CREATE INDEX IF NOT EXISTS idx_apl_machine ON assembly_production_log (machine_id);

CREATE INDEX IF NOT EXISTS idx_apl_device ON assembly_production_log (device_id);

CREATE INDEX IF NOT EXISTS idx_apl_station ON assembly_production_log (station);

CREATE INDEX IF NOT EXISTS idx_apl_prod_time ON assembly_production_log (production_time DESC);

CREATE INDEX IF NOT EXISTS idx_apl_tenant_machine_time ON assembly_production_log (
    tenant_id,
    machine_id,
    production_time DESC
);