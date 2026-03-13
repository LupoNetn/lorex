-- name: GetDelivery :one
SELECT * FROM deliveries
WHERE id = $1 LIMIT 1;

-- name: ListDeliveries :many
SELECT * FROM deliveries
ORDER BY created_at DESC;

-- name: ListDeliveriesByCustomer :many
SELECT * FROM deliveries
WHERE customer_id = $1
ORDER BY created_at DESC;

-- name: ListDeliveriesByDriver :many
SELECT * FROM deliveries
WHERE driver_id = $1
ORDER BY created_at DESC;

-- name: CreateDelivery :one
INSERT INTO deliveries (
    customer_id, pickup_address, delivery_address, pickup_lat, pickup_lng, 
    delivery_lat, delivery_lng, package_type, weight, price
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: UpdateDeliveryStatus :one
UPDATE deliveries
SET status = $2, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: AssignDriver :one
UPDATE deliveries
SET driver_id = $2, status = 'assigned', updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateDelivery :one
UPDATE deliveries
SET 
  customer_id = coalesce(sqlc.narg('customer_id'), customer_id),
  driver_id = coalesce(sqlc.narg('driver_id'), driver_id),
  pickup_address = coalesce(sqlc.narg('pickup_address'), pickup_address),
  delivery_address = coalesce(sqlc.narg('delivery_address'), delivery_address),
  pickup_lat = coalesce(sqlc.narg('pickup_lat'), pickup_lat),
  pickup_lng = coalesce(sqlc.narg('pickup_lng'), pickup_lng),
  delivery_lat = coalesce(sqlc.narg('delivery_lat'), delivery_lat),
  delivery_lng = coalesce(sqlc.narg('delivery_lng'), delivery_lng),
  status = coalesce(sqlc.narg('status'), status),
  package_type = coalesce(sqlc.narg('package_type'), package_type),
  weight = coalesce(sqlc.narg('weight'), weight),
  price = coalesce(sqlc.narg('price'), price),
  updated_at = NOW()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteDelivery :exec
DELETE FROM deliveries
WHERE id = $1;
