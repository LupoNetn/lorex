-- name: GetCustomer :one
SELECT * FROM customers
WHERE id = $1 LIMIT 1;

-- name: GetCustomerByEmail :one
SELECT * FROM customers
WHERE email = $1 LIMIT 1;

-- name: ListCustomers :many
SELECT * FROM customers
ORDER BY name;

-- name: CreateCustomer :one
INSERT INTO customers (
    company_id, name, email, phone, password, dob, gender, state_residence, country_residence, nationality
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: UpdateCustomer :one
UPDATE customers
  set company_id = coalesce(sqlc.narg('company_id'), company_id),
  name = coalesce(sqlc.narg('name'), name),
  email = coalesce(sqlc.narg('email'), email),
  phone = coalesce(sqlc.narg('phone'), phone),
  password = coalesce(sqlc.narg('password'), password),
  dob = coalesce(sqlc.narg('dob'), dob),
  gender = coalesce(sqlc.narg('gender'), gender),
  state_residence = coalesce(sqlc.narg('state_residence'), state_residence),
  country_residence = coalesce(sqlc.narg('country_residence'), country_residence),
  nationality = coalesce(sqlc.narg('nationality'), nationality),
  updated_at = NOW()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;
