-- name: InsertDevice :exec
INSERT INTO devices (id, user_id, device_id, device_name, os_version, created_at, last_used)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateLastUsed :exec
UPDATE devices
SET last_used = $2
WHERE id = $1;

-- name: UpdateDeviceName :exec
UPDATE devices
SET device_name = $2
WHERE id = $1;

-- name: GetDevicesByUserID :many
SELECT * FROM devices
WHERE user_id = $1
ORDER BY last_used DESC;

-- name: GetDeviceByID :one
SELECT * FROM devices
WHERE id = $1;

-- name: DeleteDevice :exec
DELETE FROM devices
WHERE id = $1;

-- name: DeleteDevicesByUserID :exec
DELETE FROM devices
WHERE user_id = $1;

-- name: CountDevicesByUserID :one
SELECT COUNT(*) FROM devices
WHERE user_id = $1;

-- name: GetUsersWithManyDevices :many
SELECT user_id, COUNT(*) as device_count
FROM devices
GROUP BY user_id
HAVING COUNT(*) > $1
ORDER BY device_count DESC;

-- name: GetUnusedDevices :many
SELECT * FROM devices
WHERE last_used < NOW() - INTERVAL '30 days';

-- name: DeleteOldDevices :exec
DELETE FROM devices
WHERE last_used < NOW() - INTERVAL '1 year';
