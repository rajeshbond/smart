CREATE TABLE device_master
(
    id                      BIGSERIAL PRIMARY KEY,

-- Device Identity
device_id VARCHAR(50) NOT NULL UNIQUE,
serial_number VARCHAR(100) NOT NULL UNIQUE,

-- Device Information
model VARCHAR(100) NOT NULL,
hardware_version VARCHAR(20),
firmware_version VARCHAR(20),
manufactured_at TIMESTAMP,

-- Factory Provisioning
mqtt_username           VARCHAR(100) NOT NULL UNIQUE,
    mqtt_password           VARCHAR(255) NOT NULL,

    softap_ssid             VARCHAR(100) NOT NULL,
    softap_password         VARCHAR(100) NOT NULL,

    device_secret           VARCHAR(255) NOT NULL UNIQUE,

    chip_id                 VARCHAR(100) UNIQUE,

    mac_address_wifi        VARCHAR(20),
    mac_address_ethernet    VARCHAR(20),

    communication_type      VARCHAR(20) NOT NULL DEFAULT 'WIFI',

    device_status           VARCHAR(20) NOT NULL DEFAULT 'IN_STOCK',

    last_seen_at            TIMESTAMP,

    is_active               BOOLEAN NOT NULL DEFAULT TRUE,
    is_deleted              BOOLEAN NOT NULL DEFAULT FALSE,

    notes                   TEXT,

    created_by              BIGINT,
    updated_by              BIGINT,

    created_at              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-------------------------------------------------------------------
-- Indexes
-------------------------------------------------------------------

CREATE INDEX idx_device_master_model ON device_master (model);

CREATE INDEX idx_device_master_status ON device_master (device_status);

CREATE INDEX idx_device_master_comm_type ON device_master (communication_type);

CREATE INDEX idx_device_master_last_seen ON device_master (last_seen_at);

CREATE INDEX idx_device_master_created_at ON device_master (created_at);

CREATE INDEX idx_device_master_is_active ON device_master (is_active);

CREATE INDEX idx_device_master_is_deleted ON device_master (is_deleted);

CREATE INDEX idx_device_master_hw_version ON device_master (hardware_version);

CREATE INDEX idx_device_master_fw_version ON device_master (firmware_version);