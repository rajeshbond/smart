package store

//------------------------------------------------------------------------------
// CREATE DEVICE
//------------------------------------------------------------------------------

const CreateDevice = `
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

    notes,

    created_by,
    updated_by
)
VALUES
(
    $1,
    $2,

    $3,
    $4,
    $5,
    $6,

    $7,
    $8,

    $9,
    $10,

    $11,

    $12,

    $13,
    $14,

    $15,
    $16,

    $17,

    $18,
    $19
)
RETURNING id;
`

//------------------------------------------------------------------------------
// UPDATE DEVICE
//------------------------------------------------------------------------------

const UpdateDevice = `
UPDATE device_master
SET
    model                 = $2,
    hardware_version      = $3,
    firmware_version      = $4,
    manufactured_at       = $5,

    mqtt_username         = $6,
    mqtt_password         = $7,

    softap_ssid           = $8,
    softap_password       = $9,

    device_secret         = $10,

    chip_id               = $11,

    mac_address_wifi      = $12,
    mac_address_ethernet  = $13,

    communication_type    = $14,

    device_status         = $15,

    notes                 = $16,

    is_active             = $17,

    updated_by            = $18,

    updated_at            = CURRENT_TIMESTAMP

WHERE
    id = $1
AND
    is_deleted = FALSE;
`

const deviceColumns = `
	id,
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

	mqtt_registration_status,
	mqtt_registered_at,
	mqtt_registered_by,

	is_active,
	is_deleted,

	notes,

	created_at,
	created_by,

	updated_at,
	updated_by
`

const GetDeviceByID = `
SELECT
` + deviceColumns + `
FROM
	device_master
WHERE
	id = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get Device By Device ID
//------------------------------------------------------------------------------

const GetDeviceByDeviceID = `
SELECT
` + deviceColumns + `
FROM
	device_master
WHERE
	device_id = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get Device By MQTT Username
//------------------------------------------------------------------------------

const GetDeviceByMQTTUsername = `
SELECT
` + deviceColumns + `
FROM
	device_master
WHERE
	mqtt_username = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get Device By Device Secret
//------------------------------------------------------------------------------

const GetDeviceBySecret = `
SELECT
` + deviceColumns + `
FROM
	device_master
WHERE
	device_secret = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get Device By Chip ID
//------------------------------------------------------------------------------

const GetDeviceByChipID = `
SELECT
` + deviceColumns + `
FROM
	device_master
WHERE
	chip_id = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get Device By Serial Number
//------------------------------------------------------------------------------

const GetDeviceBySerialNumber = `
SELECT
` + deviceColumns + `
FROM
	device_master
WHERE
	serial_number = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get Device By WiFi MAC Address
//------------------------------------------------------------------------------

const GetDeviceByWiFiMAC = `
SELECT
` + deviceColumns + `
FROM
	device_master
WHERE
	mac_address_wifi = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get Device By Ethernet MAC Address
//------------------------------------------------------------------------------

const GetDeviceByEthernetMAC = `
SELECT
` + deviceColumns + `
FROM
	device_master
WHERE
	mac_address_ethernet = $1
AND
	is_deleted = FALSE;
`

//==============================================================================
// EXISTS
//==============================================================================

//------------------------------------------------------------------------------
// Exists By ID
//------------------------------------------------------------------------------

const ExistsByID = `
SELECT EXISTS
(
	SELECT 1
	FROM device_master
	WHERE
		id = $1
	AND
		is_deleted = FALSE
);
`

//------------------------------------------------------------------------------
// Exists By Device ID
//------------------------------------------------------------------------------

const ExistsByDeviceID = `
SELECT EXISTS
(
	SELECT 1
	FROM device_master
	WHERE
		device_id = $1
	AND
		is_deleted = FALSE
);
`

//------------------------------------------------------------------------------
// Exists By Serial Number
//------------------------------------------------------------------------------

const ExistsBySerialNumber = `
SELECT EXISTS
(
	SELECT 1
	FROM device_master
	WHERE
		serial_number = $1
	AND
		is_deleted = FALSE
);
`

//------------------------------------------------------------------------------
// Exists By MQTT Username
//------------------------------------------------------------------------------

const ExistsByMQTTUsername = `
SELECT EXISTS
(
	SELECT 1
	FROM device_master
	WHERE
		mqtt_username = $1
	AND
		is_deleted = FALSE
);
`

//------------------------------------------------------------------------------
// Exists By Chip ID
//------------------------------------------------------------------------------

const ExistsByChipID = `
SELECT EXISTS
(
	SELECT 1
	FROM device_master
	WHERE
		chip_id = $1
	AND
		is_deleted = FALSE
);
`

//------------------------------------------------------------------------------
// Exists By Device Secret
//------------------------------------------------------------------------------

const ExistsByDeviceSecret = `
SELECT EXISTS
(
	SELECT 1
	FROM device_master
	WHERE
		device_secret = $1
	AND
		is_deleted = FALSE
);
`

//------------------------------------------------------------------------------
// Exists By WiFi MAC Address
//------------------------------------------------------------------------------

const ExistsByWiFiMAC = `
SELECT EXISTS
(
	SELECT 1
	FROM device_master
	WHERE
		mac_address_wifi = $1
	AND
		is_deleted = FALSE
);
`

//------------------------------------------------------------------------------
// Exists By Ethernet MAC Address
//------------------------------------------------------------------------------

const ExistsByEthernetMAC = `
SELECT EXISTS
(
	SELECT 1
	FROM device_master
	WHERE
		mac_address_ethernet = $1
	AND
		is_deleted = FALSE
);
`

//------------------------------------------------------------------------------
// Soft Delete Device
//------------------------------------------------------------------------------

const DeleteDevice = `
UPDATE device_master
SET
	is_deleted = TRUE,
	updated_by = $2,
	updated_at = CURRENT_TIMESTAMP
WHERE
	id = $1
AND
	is_deleted = FALSE;
`
const ListDevices = `
SELECT
` + deviceColumns + `
FROM
	device_master
WHERE
	is_deleted = FALSE
ORDER BY
	id DESC
LIMIT $1
OFFSET $2;
`
const UpdateMQTTRegistration = `
UPDATE device_master
SET
    mqtt_registration_status = $1,
    mqtt_registered_at       = $2,
    mqtt_registered_by       = $3,
    updated_by               = $4,
    updated_at               = NOW()
WHERE
    device_id = $5
    AND is_deleted = FALSE;
`
