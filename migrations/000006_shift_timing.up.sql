-- =========================================
-- SHIFT TIMING TABLE
-- =========================================

CREATE TABLE IF NOT EXISTS shift_timing (
    id SERIAL PRIMARY KEY,
    tenant_shift_id INTEGER NOT NULL,
    shift_start TIME NOT NULL,
    shift_end TIME NOT NULL,
    created_by INTEGER,
    updated_by INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_shift_tenant FOREIGN KEY (tenant_shift_id) REFERENCES tenant_shift (id) ON DELETE CASCADE,
    CONSTRAINT uq_shift_time UNIQUE (tenant_shift_id),
    CONSTRAINT chk_shift_time CHECK (shift_start <> shift_end)
);

-- =========================================
-- TRIGGER
-- =========================================

DROP TRIGGER IF EXISTS trg_update_shift_timing_updated_at ON shift_timing;

CREATE TRIGGER trg_update_shift_timing_updated_at
BEFORE UPDATE
ON shift_timing
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- =========================================
-- INDEX
-- =========================================

CREATE INDEX IF NOT EXISTS idx_shift_timing ON shift_timing (tenant_shift_id);