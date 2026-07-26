CREATE TABLE IF NOT EXISTS tenant_shift_master (
    id BIGSERIAL PRIMARY KEY,
    tenant_id UUID NOT NULL,
    shift_code BIGINT NOT NULL,
    shift_name VARCHAR(100) NOT NULL,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_by UUID,
    updated_by UUID,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_by UUID,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT uq_tenant_shift_code UNIQUE (tenant_id, shift_code),
    CONSTRAINT uq_tenant_shift_name UNIQUE (tenant_id, shift_name)
);

CREATE INDEX IF NOT EXISTS idx_tenant_shift_master_tenant ON tenant_shift_master (tenant_id);

CREATE INDEX IF NOT EXISTS idx_tenant_shift_master_active ON tenant_shift_master (tenant_id, is_active);

CREATE INDEX IF NOT EXISTS idx_tenant_shift_master_deleted ON tenant_shift_master (is_deleted);