-- =========================================
-- TENANT SHIFT TABLE
-- =========================================

CREATE TABLE IF NOT EXISTS tenant_shift (
    id SERIAL PRIMARY KEY,
    tenant_id INTEGER NOT NULL,
    shift_name VARCHAR(50) NOT NULL,
    created_by INTEGER,
    updated_by INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_tenant_shift FOREIGN KEY (tenant_id) REFERENCES tenant (id) ON DELETE CASCADE,
    CONSTRAINT uq_tenant_shift UNIQUE (tenant_id, shift_name)
);

-- =========================================
-- FUNCTION
-- =========================================

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- =========================================
-- TRIGGER
-- =========================================

DROP TRIGGER IF EXISTS trg_update_tenant_shift_updated_at ON tenant_shift;

CREATE TRIGGER trg_update_tenant_shift_updated_at
BEFORE UPDATE
ON tenant_shift
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- =========================================
-- INDEX
-- =========================================

CREATE INDEX IF NOT EXISTS idx_tenant_shift_tenant ON tenant_shift (tenant_id);