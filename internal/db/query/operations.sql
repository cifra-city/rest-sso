-- name: CreateOperation :exec
INSERT INTO operations (user_id, operation, device_data, success, failure_reason)
VALUES ($1, $2, $3, $4, $5);