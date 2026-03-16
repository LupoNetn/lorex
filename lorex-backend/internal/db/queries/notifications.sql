-- name: CreateNotification :exec
INSERT INTO notifications (receiver_id, receiver_type, company_id, message, type, read)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetNotifications :many
SELECT * FROM notifications WHERE receiver_id = $1;

-- name: MarkNotificationAsRead :exec
UPDATE notifications SET read = true WHERE id = $1;


