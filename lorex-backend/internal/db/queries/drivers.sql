-- name: GetDriver :one
SELECT * FROM drivers
WHERE id = $1 LIMIT 1;

-- name: GetDriverByEmail :one
SELECT * FROM drivers
WHERE email = $1 LIMIT 1;

-- name: ListDrivers :many
SELECT * FROM drivers
ORDER BY name;

-- name: CreateDriver :one
INSERT INTO drivers (
    company_id, name, email, phone, password, dob, gender, state_residence, country_residence, nationality
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: UpdateDriver :one
UPDATE drivers
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
  available = coalesce(sqlc.narg('available'), available),
  updated_at = NOW()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteDriver :exec
DELETE FROM drivers
WHERE id = $1;
