-- name: InsertLoginHistory :exec
INSERT INTO login_history (id, user_id, device_id, ip_address, login_time, success, failure_reason)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: GetLoginHistoryByUserID :many
SELECT * FROM login_history
WHERE user_id = $1
ORDER BY login_time DESC;

-- name: GetLoginHistoryByDeviceID :many
SELECT * FROM login_history
WHERE device_id = $1
ORDER BY login_time DESC;

-- name: GetRecentSuccessfulLogins :many
SELECT * FROM login_history
WHERE user_id = $1 AND success = true
ORDER BY login_time DESC
    LIMIT $2;

-- name: DeleteLoginHistoryByUserID :exec
DELETE FROM login_history
WHERE user_id = $1;

-- name: DeleteLoginHistoryByDeviceID :exec
DELETE FROM login_history
WHERE device_id = $1;

-- name: CountSuccessfulLogins :one
SELECT COUNT(*) FROM login_history
WHERE user_id = $1 AND success = true;

-- name: CountFailedLogins :one
SELECT COUNT(*) FROM login_history
WHERE user_id = $1 AND success = false;

-- name: GetUsersWithFrequentFailedLogins :many
SELECT user_id, COUNT(*) as failure_count
FROM login_history
WHERE success = false
GROUP BY user_id
HAVING COUNT(*) > $1
ORDER BY failure_count DESC;

-- name: DeleteOldLoginHistory :exec
DELETE FROM login_history
WHERE login_time < NOW() - INTERVAL '90 days';
