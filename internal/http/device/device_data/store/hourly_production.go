package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

// ============================================================
// GET CURRENT SHIFT
// ============================================================
//
// Handles:
//
// Normal shift:
//     08:00 -> 20:00
//
// Overnight shift:
//     20:00 -> 08:00
//
// ============================================================

func (s *Store) GetCurrentShift(
	ctx context.Context,
	tx *sql.Tx,
	tenantID string,
	currentTime time.Time,
) (*dto.ShiftInfoHR, error) {

	const query = `
		SELECT
			ts.id,
			st.id,
			ts.shift_name,
			st.shift_start,
			st.shift_end

		FROM public.tenant_shift ts

		INNER JOIN public.shift_timing st
			ON st.tenant_shift_id = ts.id

		WHERE ts.tenant_id = $1

		AND
		(
			(
				-- Normal shift
				st.shift_start <= st.shift_end
				AND $2::time >= st.shift_start
				AND $2::time < st.shift_end
			)

			OR

			(
				-- Overnight shift
				st.shift_start > st.shift_end

				AND
				(
					$2::time >= st.shift_start
					OR
					$2::time < st.shift_end
				)
			)
		)

		ORDER BY st.shift_start

		LIMIT 1
	`

	var item dto.ShiftInfoHR

	err := tx.QueryRowContext(
		ctx,
		query,
		tenantID,
		currentTime,
	).Scan(
		&item.TenantShiftID,
		&item.ShiftTimingID,
		&item.ShiftName,
		&item.ShiftStart,
		&item.ShiftEnd,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(
				"current shift not found",
			)
		}

		return nil, fmt.Errorf(
			"get current shift: %w",
			err,
		)
	}

	return &item, nil
}

// ============================================================
// GET PREVIOUS COMPLETED HOUR SLOT
// ============================================================
//
// Example:
//
// Current time = 22:03
//
// Returns:
//
// 21:00 -> 22:00
//
// At:
//
// 08:03
//
// Returns:
//
// 07:00 -> 08:00
//
// ============================================================

func (s *Store) GetPreviousHourSlot(
	ctx context.Context,
	tx *sql.Tx,
	tenantID string,
	shiftTimingID int64,
	currentTime time.Time,
) (*dto.HourSlot, error) {

	const query = `
		SELECT
			id,
			shift_timing_id,
			slot_index,
			slot_start,
			slot_end

		FROM public.shift_hour_slot

		WHERE tenant_id = $1

		  AND shift_timing_id = $2

		  AND slot_end <= $3::time

		ORDER BY slot_end DESC

		LIMIT 1
	`

	var item dto.HourSlot

	err := tx.QueryRowContext(
		ctx,
		query,
		tenantID,
		shiftTimingID,
		currentTime,
	).Scan(
		&item.ShiftHourSlotID,
		&item.ShiftTimingID,
		&item.SlotIndex,
		&item.SlotStart,
		&item.SlotEnd,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(
				"previous completed hour slot not found",
			)
		}

		return nil, fmt.Errorf(
			"get previous hour slot: %w",
			err,
		)
	}

	return &item, nil
}

// ============================================================
// GET PRODUCTION STATISTICS
// ============================================================
//
// # IMPORTANT FORMULA
//
// Previous hour FIRST count
//
//	|
//	|
//	v
//
// Current hour FIRST count
//
//	|
//	|
//	v
//
// # Current First - Previous First
//
// ============================================================
//
// Example:
//
// 08:00 first count = 1000
// 09:00 first count = 1020
// 10:00 first count = 1050
//
// 08-09:
//
//	1020 - 1000 = 20
//
// 09-10:
//
//	1050 - 1020 = 30
//
// ============================================================
//
// IMPORTANT:
//
// We DO NOT use:
//
//	current hour LAST count
//	-
//	current hour FIRST count
//
// anymore.
//
// ============================================================
func (s *Store) GetProductionStats(
	ctx context.Context,
	tx *sql.Tx,

	tenantID string,
	deviceID string,
	station string,

	currentShiftStart time.Time,

	actualStart time.Time,
	actualEnd time.Time,
) (*dto.HourlyProductionStats, error) {

	const query = `
		WITH

		-- ======================================================
		-- CURRENT HOUR FIRST COUNT
		-- ======================================================
		--
		-- Example:
		--
		-- Current hour:
		-- 12:00 -> 13:00
		--
		-- Production:
		--
		-- 12:05 -> 7497
		-- 12:20 -> 7505
		-- 12:40 -> 7510
		--
		-- Result:
		--
		-- 7497
		--
		-- ======================================================

		current_hour_first AS
		(
			SELECT
				apl.production_count

			FROM public.assembly_production_log apl

			WHERE apl.tenant_id = $1
			  AND apl.device_id = $2
			  AND ($3 = '' OR apl.station = $3)

			  AND apl.production_time >= $5
			  AND apl.production_time < $6

			ORDER BY
				apl.production_time ASC,
				apl.id ASC

			LIMIT 1
		),

		-- ======================================================
		-- PREVIOUS SHIFT LAST COUNT
		-- ======================================================
		--
		-- Get the LAST cumulative production count before the
		-- current shift started.
		--
		-- Example:
		--
		-- Previous shift:
		--
		-- 19:50 -> 7490
		-- 19:58 -> 7497
		--
		-- Current shift:
		--
		-- 20:00
		--
		-- Result:
		--
		-- 7497
		--
		-- ======================================================

		previous_shift_last AS
		(
			SELECT
				apl.production_count

			FROM public.assembly_production_log apl

			WHERE apl.tenant_id = $1
			  AND apl.device_id = $2
			  AND ($3 = '' OR apl.station = $3)

			  AND apl.production_time < $4

			ORDER BY
				apl.production_time DESC,
				apl.id DESC

			LIMIT 1
		),

		-- ======================================================
		-- LAST COUNT BEFORE CURRENT HOUR
		-- ======================================================
		--
		-- Used when:
		--
		-- 1. Current hour has no production
		-- 2. Current hour is NOT the first hour of shift
		--
		-- Example:
		--
		-- 11:55 -> 7510
		--
		-- 12:00 -> 13:00
		-- NO PRODUCTION
		--
		-- Result:
		--
		-- 7510
		--
		-- Fallback start:
		--
		-- 7510 + 1 = 7511
		--
		-- ======================================================

		last_before_current_hour AS
		(
			SELECT
				apl.production_count

			FROM public.assembly_production_log apl

			WHERE apl.tenant_id = $1
			  AND apl.device_id = $2
			  AND ($3 = '' OR apl.station = $3)

			  AND apl.production_time < $5

			ORDER BY
				apl.production_time DESC,
				apl.id DESC

			LIMIT 1
		),

		-- ======================================================
		-- NEXT HOUR FIRST COUNT
		-- ======================================================
		--
		-- Example:
		--
		-- Current hour:
		-- 12:00 -> 13:00
		--
		-- Next hour:
		-- 13:00 -> 14:00
		--
		-- First production:
		--
		-- 13:02 -> 7520
		--
		-- Result:
		--
		-- 7520
		--
		-- ======================================================

		next_hour_first AS
		(
			SELECT
				apl.production_count

			FROM public.assembly_production_log apl

			WHERE apl.tenant_id = $1
			  AND apl.device_id = $2
			  AND ($3 = '' OR apl.station = $3)

			  AND apl.production_time >= $6
			  AND apl.production_time < ($6 + INTERVAL '1 hour')

			ORDER BY
				apl.production_time ASC,
				apl.id ASC

			LIMIT 1
		),

		-- ======================================================
		-- RESOLVED START COUNT
		-- ======================================================
		--
		-- Priority:
		--
		-- 1. Current hour first count
		--
		-- 2. First hour of shift:
		--       previous shift last + 1
		--
		-- 3. Normal hour:
		--       last count before current hour + 1
		--
		-- 4. No baseline:
		--       0
		--
		-- ======================================================

		resolved_start AS
		(
			SELECT

				CASE

					-- ==========================================
					-- CASE 1
					--
					-- Current hour has production.
					-- ==========================================

					WHEN
						(
							SELECT production_count
							FROM current_hour_first
						) IS NOT NULL

					THEN
						(
							SELECT production_count
							FROM current_hour_first
						)

					-- ==========================================
					-- CASE 2
					--
					-- No current production.
					--
					-- Current hour is the FIRST hour of shift.
					--
					-- Use:
					--
					-- previous shift last + 1
					-- ==========================================

					WHEN $5 = $4

					THEN
						COALESCE(
							(
								SELECT production_count
								FROM previous_shift_last
							) + 1,
							0
						)

					-- ==========================================
					-- CASE 3
					--
					-- No current production.
					--
					-- Not the first hour of shift.
					--
					-- Use:
					--
					-- last count before current hour + 1
					-- ==========================================

					WHEN
						(
							SELECT production_count
							FROM last_before_current_hour
						) IS NOT NULL

					THEN
						(
							SELECT production_count
							FROM last_before_current_hour
						) + 1

					-- ==========================================
					-- CASE 4
					--
					-- No baseline available.
					-- ==========================================

					ELSE 0

				END AS start_production_count
		),

		-- ======================================================
		-- CYCLE STATISTICS
		-- ======================================================
		--
		-- Only calculate cycle statistics for the current hour.
		--
		-- ======================================================

		cycle_stats AS
		(
			SELECT

				COUNT(cycle_time_sec)
					FILTER (
						WHERE cycle_time_sec IS NOT NULL
					) AS cycle_count,

				MIN(cycle_time_sec)
					FILTER (
						WHERE cycle_time_sec IS NOT NULL
						AND cycle_time_sec > 0
					) AS min_cycle_time_sec,

				AVG(cycle_time_sec)
					FILTER (
						WHERE cycle_time_sec IS NOT NULL
						AND cycle_time_sec > 0
					) AS avg_cycle_time_sec,

				MAX(cycle_time_sec)
					FILTER (
						WHERE cycle_time_sec IS NOT NULL
						AND cycle_time_sec > 0
					) AS max_cycle_time_sec

			FROM public.assembly_production_log apl

			WHERE apl.tenant_id = $1
			  AND apl.device_id = $2
			  AND ($3 = '' OR apl.station = $3)

			  AND apl.production_time >= $5
			  AND apl.production_time < $6
		)

		-- ======================================================
		-- FINAL RESULT
		-- ======================================================

		SELECT

			-- ==================================================
			-- START PRODUCTION COUNT
			-- ==================================================

			resolved_start.start_production_count
				AS start_production_count,

			-- ==================================================
			-- END PRODUCTION COUNT
			-- ==================================================
			--
			-- First production count of NEXT hour.
			--
			-- ==================================================

			(
				SELECT production_count
				FROM next_hour_first
			) AS end_production_count,

			-- ==================================================
			-- HOURLY PRODUCTION COUNT
			-- ==================================================
			--
			-- Formula:
			--
			-- next hour first
			-- -
			-- resolved start
			--
			-- Example:
			--
			-- Start = 7497
			-- End   = 7520
			--
			-- Result = 23
			--
			-- ==================================================

			CASE

				-- ----------------------------------------------
				-- No next-hour production yet.
				--
				-- Cannot finalize the hourly production.
				-- ----------------------------------------------

				WHEN
					(
						SELECT production_count
						FROM next_hour_first
					) IS NULL

				THEN 0

				-- ----------------------------------------------
				-- No starting baseline.
				-- ----------------------------------------------

				WHEN
					resolved_start.start_production_count = 0

				THEN 0

				-- ----------------------------------------------
				-- Normal calculation.
				-- ----------------------------------------------

				ELSE GREATEST(

					(
						SELECT production_count
						FROM next_hour_first
					)

					-

					resolved_start.start_production_count,

					0
				)

			END AS hourly_production_count,

			-- ==================================================
			-- CYCLE COUNT
			-- ==================================================

			COALESCE(
				cycle_stats.cycle_count,
				0
			) AS cycle_count,

			-- ==================================================
			-- MIN CYCLE TIME
			-- ==================================================

			COALESCE(
				cycle_stats.min_cycle_time_sec,
				0
			) AS min_cycle_time_sec,

			-- ==================================================
			-- AVG CYCLE TIME
			-- ==================================================

			COALESCE(
				cycle_stats.avg_cycle_time_sec,
				0
			) AS avg_cycle_time_sec,

			-- ==================================================
			-- MAX CYCLE TIME
			-- ==================================================

			COALESCE(
				cycle_stats.max_cycle_time_sec,
				0
			) AS max_cycle_time_sec

		FROM
			resolved_start,
			cycle_stats
	`

	var item dto.HourlyProductionStats

	err := tx.QueryRowContext(
		ctx,
		query,

		tenantID,
		deviceID,
		station,

		currentShiftStart,
		actualStart,
		actualEnd,
	).Scan(
		&item.StartProductionCount,
		&item.EndProductionCount,

		&item.HourlyProductionCount,

		&item.CycleCount,

		&item.MinCycleTimeSec,
		&item.AvgCycleTimeSec,
		&item.MaxCycleTimeSec,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"get production statistics: %w",
			err,
		)
	}

	return &item, nil
}

// ============================================================
// SAVE HOURLY PRODUCTION
// ============================================================

func (s *Store) SaveHourlyProduction(
	ctx context.Context,
	tx *sql.Tx,
	item *dto.HourlyProduction,
) error {

	const query = `
		INSERT INTO public.assembly_hourly_production
		(
			tenant_id,
			customer_id,
			device_id,
			machine_id,
			station,
			variant,

			production_day,

			tenant_shift_id,
			shift_timing_id,
			shift_hour_slot_id,
			slot_index,

			slot_start,
			slot_end,

			actual_start,
			actual_end,

			start_production_count,
			end_production_count,
			hourly_production_count,

			cycle_count,
			min_cycle_time_sec,
			avg_cycle_time_sec,
			max_cycle_time_sec,

			created_at,
			updated_at
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

			$19,
			$20,
			$21,
			$22,

			NOW(),
			NOW()
		)

		ON CONFLICT
		(
			tenant_id,
			device_id,
			machine_id,
			station,
			production_day,
			shift_timing_id,
			shift_hour_slot_id
		)

		DO UPDATE SET

			customer_id =
				EXCLUDED.customer_id,

			variant =
				EXCLUDED.variant,

			actual_start =
				EXCLUDED.actual_start,

			actual_end =
				EXCLUDED.actual_end,

			start_production_count =
				EXCLUDED.start_production_count,

			end_production_count =
				EXCLUDED.end_production_count,

			hourly_production_count =
				EXCLUDED.hourly_production_count,

			cycle_count =
				EXCLUDED.cycle_count,

			min_cycle_time_sec =
				EXCLUDED.min_cycle_time_sec,

			avg_cycle_time_sec =
				EXCLUDED.avg_cycle_time_sec,

			max_cycle_time_sec =
				EXCLUDED.max_cycle_time_sec,

			updated_at =
				NOW()
	`

	_, err := tx.ExecContext(
		ctx,
		query,

		item.TenantID,
		item.CustomerID,
		item.DeviceID,
		item.MachineID,
		item.Station,
		item.Variant,

		item.ProductionDay,

		item.TenantShiftID,
		item.ShiftTimingID,
		item.ShiftHourSlotID,
		item.SlotIndex,

		item.SlotStart,
		item.SlotEnd,

		item.ActualStart,
		item.ActualEnd,

		item.StartProductionCount,
		item.EndProductionCount,
		item.HourlyProductionCount,

		item.CycleCount,
		item.MinCycleTimeSec,
		item.AvgCycleTimeSec,
		item.MaxCycleTimeSec,
	)

	if err != nil {
		return fmt.Errorf(
			"save hourly production: %w",
			err,
		)
	}

	return nil
}
