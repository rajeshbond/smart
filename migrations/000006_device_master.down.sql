DROP INDEX IF EXISTS idx_device_master_fw_version;

DROP INDEX IF EXISTS idx_device_master_hw_version;

DROP INDEX IF EXISTS idx_device_master_is_deleted;

DROP INDEX IF EXISTS idx_device_master_is_active;

DROP INDEX IF EXISTS idx_device_master_created_at;

DROP INDEX IF EXISTS idx_device_master_last_seen;

DROP INDEX IF EXISTS idx_device_master_comm_type;

DROP INDEX IF EXISTS idx_device_master_status;

DROP INDEX IF EXISTS idx_device_master_model;

DROP TABLE IF EXISTS device_master;