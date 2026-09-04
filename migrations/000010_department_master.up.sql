CREATE TABLE IF NOT EXISTS department_master (
    id SERIAL PRIMARY KEY,
    department VARCHAR(100) NOT NULL UNIQUE,
    created_by INTEGER,
    updated_by INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMPTZ
);