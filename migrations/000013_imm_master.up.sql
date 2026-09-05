
-- ============================================================
-- CREATE TABLE: imm
-- Injection Moulding Machine Master
-- ============================================================

CREATE TABLE IF NOT EXISTS imm (
    -- ========================================================
    -- PRIMARY KEY
    -- ========================================================

    sr BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    -- ========================================================
    -- TENANT / DEPARTMENT
    -- ========================================================

    tenant_id INT NOT NULL,
    dept_id   INT NOT NULL,

    -- ========================================================
    -- MACHINE DETAILS
    -- ========================================================

    machine_name VARCHAR(100) NOT NULL,
    machine_no   VARCHAR(100) NOT NULL,
    machine_make VARCHAR(100),

    -- ========================================================
    -- MACHINE SPECIFICATIONS
    -- ========================================================

    tie_bar_distance   FLOAT,
    platten_size       FLOAT,
    location_ring_size FLOAT,

    -- ========================================================
    -- AUDIT USER
    -- ========================================================

    created_by INT,
    updated_by INT,

    -- ========================================================
    -- SOFT DELETE
    -- ========================================================

    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_by INT,

    -- ========================================================
    -- TIMESTAMPS
    -- ========================================================

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,

    -- ========================================================
    -- FOREIGN KEY: TENANT
    -- ========================================================

    CONSTRAINT fk_imm_tenant
        FOREIGN KEY (tenant_id)
        REFERENCES tenant(id),

    -- ========================================================
    -- FOREIGN KEY: DEPARTMENT
    -- ========================================================

    CONSTRAINT fk_imm_department
        FOREIGN KEY (dept_id)
        REFERENCES department_master(id)
);

-- ============================================================
-- UNIQUE MACHINE NUMBER PER TENANT
--
-- Same machine_no is allowed in different tenants.
-- Same machine_no is NOT allowed twice in the same tenant.
--
-- Deleted machines are excluded from uniqueness.
-- ============================================================

CREATE UNIQUE INDEX uk_imm_tenant_machine_no
    ON imm (tenant_id, machine_no)
    WHERE is_deleted = FALSE;

-- ============================================================
-- INDEXES
-- ============================================================

CREATE INDEX idx_imm_tenant_id
    ON imm (tenant_id);

CREATE INDEX idx_imm_dept_id
    ON imm (dept_id);

CREATE INDEX idx_imm_machine_no
    ON imm (machine_no);

CREATE INDEX idx_imm_is_deleted
    ON imm (is_deleted);

CREATE INDEX idx_imm_tenant_active
    ON imm (tenant_id, is_deleted);
