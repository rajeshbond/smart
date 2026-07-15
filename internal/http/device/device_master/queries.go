package devicemaster

const (

	//------------------------------------------------------------------
	// Insert
	//------------------------------------------------------------------

	InsertDevice = `
INSERT INTO device_master
(
	device_id,
	serial_number,
	model,
	hardware_version,
	firmware_version,
	manufactured_at,
	mqtt_username,
	mqtt_password,
	softap_ssid,
	softap_password,
	device_secret,
	chip_id,
	mac_address_wifi,
	mac_address_ethernet,
	communication_type,
	device_status,
	last_seen_at,
	is_active,
	is_deleted,
	notes,
	created_by,
	updated_by
)
VALUES
(
	:device_id,
	:serial_number,
	:model,
	:hardware_version,
	:firmware_version,
	:manufactured_at,
	:mqtt_username,
	:mqtt_password,
	:softap_ssid,
	:softap_password,
	:device_secret,
	:chip_id,
	:mac_address_wifi,
	:mac_address_ethernet,
	:communication_type,
	:device_status,
	:last_seen_at,
	:is_active,
	:is_deleted,
	:notes,
	:created_by,
	:updated_by
)
RETURNING id;
`

	//------------------------------------------------------------------
	// Get By ID
	//------------------------------------------------------------------

	GetDeviceByID = `
SELECT *
FROM device_master
WHERE id=$1
AND is_deleted=false;
`

	//------------------------------------------------------------------
	// Get By Device ID
	//------------------------------------------------------------------

	GetDeviceByDeviceID = `
SELECT *
FROM device_master
WHERE device_id=$1
AND is_deleted=false;
`

	//------------------------------------------------------------------
	// Get By Serial Number
	//------------------------------------------------------------------

	GetDeviceBySerialNumber = `
SELECT *
FROM device_master
WHERE serial_number=$1
AND is_deleted=false;
`

	//------------------------------------------------------------------
	// Get By MQTT Username
	//------------------------------------------------------------------

	GetDeviceByMQTTUsername = `
SELECT *
FROM device_master
WHERE mqtt_username=$1
AND is_deleted=false;
`

	//------------------------------------------------------------------
	// Update
	//------------------------------------------------------------------

	UpdateDevice = `
UPDATE device_master
SET

	model=:model,

	hardware_version=:hardware_version,

	firmware_version=:firmware_version,

	mqtt_username=:mqtt_username,

	mqtt_password=:mqtt_password,

	softap_ssid=:softap_ssid,

	softap_password=:softap_password,

	device_secret=:device_secret,

	chip_id=:chip_id,

	mac_address_wifi=:mac_address_wifi,

	mac_address_ethernet=:mac_address_ethernet,

	communication_type=:communication_type,

	device_status=:device_status,

	notes=:notes,

	is_active=:is_active,

	updated_by=:updated_by,

	updated_at=NOW()

WHERE id=:id
AND is_deleted=false;
`

	//------------------------------------------------------------------
	// Soft Delete
	//------------------------------------------------------------------

	DeleteDevice = `
UPDATE device_master
SET

	is_deleted=true,

	updated_by=$2,

	updated_at=NOW()

WHERE id=$1;
`

	//------------------------------------------------------------------
	// Update Firmware
	//------------------------------------------------------------------

	UpdateFirmware = `
UPDATE device_master
SET

	firmware_version=$2,

	updated_at=NOW()

WHERE id=$1;
`

	//------------------------------------------------------------------
	// Update Last Seen
	//------------------------------------------------------------------

	UpdateLastSeen = `
UPDATE device_master
SET

	last_seen_at=NOW()

WHERE device_id=$1;
`

	//------------------------------------------------------------------
	// Update Status
	//------------------------------------------------------------------

	UpdateStatus = `
UPDATE device_master
SET

	device_status=$2,

	updated_at=NOW()

WHERE id=$1;
`

	//------------------------------------------------------------------
	// List
	//------------------------------------------------------------------

	ListDevices = `
SELECT *

FROM device_master

WHERE is_deleted=false
`

	//------------------------------------------------------------------
	// Count
	//------------------------------------------------------------------

	CountDevices = `
SELECT COUNT(*)

FROM device_master

WHERE is_deleted=false;
`
)
