-- name: InsertOperationHistory :exec
INSERT INTO operation_history (id, user_id, device_data, operation, success, failure_reason, ip_address)
VALUES ($1, $2, $3, $4, $5, $6, $7);