/******************************************************************************
 *
 * MODULE      : Device Master
 * FILE        : queries.go
 *
 * DESCRIPTION :
 * PostgreSQL Queries
 *
 ******************************************************************************/

package store

// -----------------------------------------------------------------------------
// INSERT
// -----------------------------------------------------------------------------

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

    :notes,

    :created_by,
    :updated_by
)
RETURNING id;
`

// -----------------------------------------------------------------------------
// UPDATE
// -----------------------------------------------------------------------------

const UpdateDevice = `
UPDATE device_master
SET

    model               = :model,
    hardware_version    = :hardware_version,
    firmware_version    = :firmware_version,
    manufactured_at     = :manufactured_at,

    mqtt_username       = :mqtt_username,
    mqtt_password       = :mqtt_password,

    softap_ssid         = :softap_ssid,
    softap_password     = :softap_password,

    device_secret       = :device_secret,

    chip_id             = :chip_id,

    mac_address_wifi    = :mac_address_wifi,
    mac_address_ethernet= :mac_address_ethernet,

    communication_type  = :communication_type,

    device_status       = :device_status,

    notes               = :notes,

    is_active           = :is_active,

    updated_by          = :updated_by,

    updated_at          = CURRENT_TIMESTAMP

WHERE
    id = :id
AND
    is_deleted = FALSE;
`

// -----------------------------------------------------------------------------
// SOFT DELETE
// -----------------------------------------------------------------------------

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

// -----------------------------------------------------------------------------
// EXISTS
// -----------------------------------------------------------------------------

const ExistsByMQTTUsername = `
SELECT EXISTS
(
    SELECT 1
    FROM device_master
    WHERE mqtt_username = $1
    AND is_deleted = FALSE
);
`

const ExistsByChipID = `
SELECT EXISTS
(
    SELECT 1
    FROM device_master
    WHERE chip_id = $1
    AND is_deleted = FALSE
);
`

//==============================================================================
// GET
//==============================================================================

//------------------------------------------------------------------------------
// Get By ID
//------------------------------------------------------------------------------

const GetDeviceByID = `
SELECT *
FROM device_master
WHERE
	id = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get By Device ID
//------------------------------------------------------------------------------

const GetDeviceByDeviceID = `
SELECT *
FROM device_master
WHERE
	device_id = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get By Serial Number
//------------------------------------------------------------------------------

const GetDeviceBySerialNumber = `
SELECT *
FROM device_master
WHERE
	serial_number = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get By MQTT Username
//------------------------------------------------------------------------------

const GetDeviceByMQTTUsername = `
SELECT *
FROM device_master
WHERE
	mqtt_username = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get By Chip ID
//------------------------------------------------------------------------------

const GetDeviceByChipID = `
SELECT *
FROM device_master
WHERE
	chip_id = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get By Device Secret
//------------------------------------------------------------------------------

const GetDeviceBySecret = `
SELECT *
FROM device_master
WHERE
	device_secret = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get By WiFi MAC Address
//------------------------------------------------------------------------------

const GetDeviceByWiFiMAC = `
SELECT *
FROM device_master
WHERE
	mac_address_wifi = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get By Ethernet MAC Address
//------------------------------------------------------------------------------

const GetDeviceByEthernetMAC = `
SELECT *
FROM device_master
WHERE
	mac_address_ethernet = $1
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get Active Devices
//------------------------------------------------------------------------------

const GetActiveDevices = `
SELECT *
FROM device_master
WHERE
	is_active = TRUE
AND
	is_deleted = FALSE
ORDER BY
	device_id;
`

//------------------------------------------------------------------------------
// Get Inactive Devices
//------------------------------------------------------------------------------

const GetInactiveDevices = `
SELECT *
FROM device_master
WHERE
	is_active = FALSE
AND
	is_deleted = FALSE
ORDER BY
	device_id;
`

//------------------------------------------------------------------------------
// Get Online Devices
//------------------------------------------------------------------------------

const GetOnlineDevices = `
SELECT *
FROM device_master
WHERE
	device_status = 'ONLINE'
AND
	is_deleted = FALSE
ORDER BY
	last_seen_at DESC;
`

//------------------------------------------------------------------------------
// Get Offline Devices
//------------------------------------------------------------------------------

const GetOfflineDevices = `
SELECT *
FROM device_master
WHERE
	device_status = 'OFFLINE'
AND
	is_deleted = FALSE
ORDER BY
	last_seen_at DESC;
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
// Exists By WiFi MAC
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
// Exists By Ethernet MAC
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

//==============================================================================
// LIST
//==============================================================================

const ListDevices = `
SELECT
	*
FROM
	device_master
`

//==============================================================================
// COUNT
//==============================================================================

const CountDevices = `
SELECT
	COUNT(*)
FROM
	device_master
`

//==============================================================================
// DEVICE PROVISIONING
//==============================================================================

//------------------------------------------------------------------------------
// Factory Provision Device
//------------------------------------------------------------------------------

const ProvisionDevice = `
UPDATE device_master
SET

	chip_id = $2,

	mac_address_wifi = $3,

	mac_address_ethernet = $4,

	updated_by = $5,

	updated_at = CURRENT_TIMESTAMP

WHERE

	id = $1

AND

	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Update WiFi MAC Address
//------------------------------------------------------------------------------

const UpdateWiFiMAC = `
UPDATE device_master
SET

	mac_address_wifi = $2,

	updated_by = $3,

	updated_at = CURRENT_TIMESTAMP

WHERE

	id = $1

AND

	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Update Ethernet MAC Address
//------------------------------------------------------------------------------

const UpdateEthernetMAC = `
UPDATE device_master
SET

	mac_address_ethernet = $2,

	updated_by = $3,

	updated_at = CURRENT_TIMESTAMP

WHERE

	id = $1

AND

	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Update Chip ID
//------------------------------------------------------------------------------

const UpdateChipID = `
UPDATE device_master
SET

	chip_id = $2,

	updated_by = $3,

	updated_at = CURRENT_TIMESTAMP

WHERE

	id = $1

AND

	is_deleted = FALSE;
`

//==============================================================================
// MQTT AUTHENTICATION
//==============================================================================

//------------------------------------------------------------------------------
// Authenticate Device
//------------------------------------------------------------------------------

const AuthenticateDevice = `
SELECT *
FROM device_master
WHERE

	device_id = $1

AND

	mqtt_username = $2

AND

	mqtt_password = $3

AND

	is_active = TRUE

AND

	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Get Device By MQTT Username
//------------------------------------------------------------------------------

const GetDeviceForMQTT = `
SELECT *
FROM device_master
WHERE

	mqtt_username = $1

AND

	is_active = TRUE

AND

	is_deleted = FALSE;
`

//==============================================================================
// COMMUNICATION
//==============================================================================

//------------------------------------------------------------------------------
// Update Communication Type
//------------------------------------------------------------------------------

const UpdateCommunicationType = `
UPDATE device_master
SET

	communication_type = $2,

	updated_by = $3,

	updated_at = CURRENT_TIMESTAMP

WHERE

	id = $1

AND

	is_deleted = FALSE;
`

//==============================================================================
// DEVICE ACTIVATION
//==============================================================================

//------------------------------------------------------------------------------
// Activate Device
//------------------------------------------------------------------------------

const ActivateDevice = `
UPDATE device_master
SET

	is_active = TRUE,

	updated_by = $2,

	updated_at = CURRENT_TIMESTAMP

WHERE

	id = $1

AND

	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Deactivate Device
//------------------------------------------------------------------------------

const DeactivateDevice = `
UPDATE device_master
SET

	is_active = FALSE,

	updated_by = $2,

	updated_at = CURRENT_TIMESTAMP

WHERE

	id = $1

AND

	is_deleted = FALSE;
`

//==============================================================================
// DEVICE SECRET
//==============================================================================

//------------------------------------------------------------------------------
// Update Device Secret
//------------------------------------------------------------------------------

const UpdateDeviceSecret = `
UPDATE device_master
SET

	device_secret = $2,

	updated_by = $3,

	updated_at = CURRENT_TIMESTAMP

WHERE

	id = $1

AND

	is_deleted = FALSE;
`

//==============================================================================
// DASHBOARD
//==============================================================================

//------------------------------------------------------------------------------
// Total Devices
//------------------------------------------------------------------------------

const CountTotalDevices = `
SELECT COUNT(*)
FROM device_master
WHERE
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Active Devices
//------------------------------------------------------------------------------

const CountActiveDevices = `
SELECT COUNT(*)
FROM device_master
WHERE
	is_active = TRUE
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Inactive Devices
//------------------------------------------------------------------------------

const CountInactiveDevices = `
SELECT COUNT(*)
FROM device_master
WHERE
	is_active = FALSE
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Online Devices
//------------------------------------------------------------------------------

const CountOnlineDevices = `
SELECT COUNT(*)
FROM device_master
WHERE
	device_status = 'ONLINE'
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Offline Devices
//------------------------------------------------------------------------------

const CountOfflineDevices = `
SELECT COUNT(*)
FROM device_master
WHERE
	device_status = 'OFFLINE'
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// In Stock
//------------------------------------------------------------------------------

const CountInStockDevices = `
SELECT COUNT(*)
FROM device_master
WHERE
	device_status = 'IN_STOCK'
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Installed
//------------------------------------------------------------------------------

const CountInstalledDevices = `
SELECT COUNT(*)
FROM device_master
WHERE
	device_status = 'INSTALLED'
AND
	is_deleted = FALSE;
`

//------------------------------------------------------------------------------
// Faulty
//------------------------------------------------------------------------------

const CountFaultyDevices = `
SELECT COUNT(*)
FROM device_master
WHERE
	device_status = 'FAULTY'
AND
	is_deleted = FALSE;
`

//==============================================================================
// HEARTBEAT
//==============================================================================

//------------------------------------------------------------------------------
// Devices Not Seen Since
//------------------------------------------------------------------------------

const GetDevicesNotSeenSince = `
SELECT *
FROM device_master
WHERE
	last_seen_at < $1
AND
	is_deleted = FALSE
ORDER BY
	last_seen_at ASC;
`

//------------------------------------------------------------------------------
// Recently Seen Devices
//------------------------------------------------------------------------------

const GetRecentlySeenDevices = `
SELECT *
FROM device_master
WHERE
	last_seen_at IS NOT NULL
AND
	is_deleted = FALSE
ORDER BY
	last_seen_at DESC;
`

//==============================================================================
// FACTORY
//==============================================================================

//------------------------------------------------------------------------------
// Devices Waiting For Provisioning
//------------------------------------------------------------------------------

const GetUnProvisionedDevices = `
SELECT *
FROM device_master
WHERE
	chip_id IS NULL
AND
	is_deleted = FALSE
ORDER BY
	created_at ASC;
`

//------------------------------------------------------------------------------
// Provisioned Devices
//------------------------------------------------------------------------------

const GetProvisionedDevices = `
SELECT *
FROM device_master
WHERE
	chip_id IS NOT NULL
AND
	is_deleted = FALSE
ORDER BY
	created_at DESC;
`
const UpdateLastSeen = `
UPDATE device_master
SET
	last_seen_at = CURRENT_TIMESTAMP,
	updated_at = CURRENT_TIMESTAMP
WHERE
	device_id = $1
AND
	is_deleted = FALSE;
`
const UpdateDeviceStatus = `
UPDATE device_master
SET
	device_status = $2,
	updated_by = $3,
	updated_at = CURRENT_TIMESTAMP
WHERE
	id = $1
	AND is_deleted = FALSE;
`

const UpdateFirmwareVersion = `
UPDATE device_master
SET
	firmware_version = $2,
	updated_by = $3,
	updated_at = CURRENT_TIMESTAMP
WHERE
	id = $1
	AND is_deleted = FALSE;
`
