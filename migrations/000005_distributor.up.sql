CREATE TABLE distributor (
    id BIGSERIAL PRIMARY KEY,
    distributor_code VARCHAR(50) NOT NULL UNIQUE,
    distributor_name VARCHAR(200) NOT NULL,
    contact_person_name VARCHAR(200),
    contact_phone VARCHAR(20),
    contact_email VARCHAR(255),
    address TEXT,
    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_distributor_code ON distributor (distributor_code);

CREATE INDEX idx_distributor_name ON distributor (distributor_name);

CREATE INDEX idx_distributor_active ON distributor (is_active);

CREATE INDEX idx_distributor_deleted ON distributor (is_deleted);