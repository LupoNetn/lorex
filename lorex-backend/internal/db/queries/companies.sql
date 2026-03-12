-- name: GetCompany :one
SELECT * FROM companies
WHERE id = $1 LIMIT 1;

-- name: GetCompanyByEmail :one
SELECT * FROM companies
WHERE email = $1 LIMIT 1;

-- name: GetCompanyBySignupCode :one
SELECT * FROM companies
WHERE customer_signup_code = $1 LIMIT 1;

-- name: ListCompanies :many
SELECT * FROM companies
ORDER BY name;

-- name: CreateCompany :one
INSERT INTO companies (
    name, description, logo, industry, phone, email, location, password, customer_signup_code
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateCompany :one
UPDATE companies
  set name = coalesce(sqlc.narg('name'), name),
  description = coalesce(sqlc.narg('description'), description),
  logo = coalesce(sqlc.narg('logo'), logo),
  industry = coalesce(sqlc.narg('industry'), industry),
  phone = coalesce(sqlc.narg('phone'), phone),
  email = coalesce(sqlc.narg('email'), email),
  location = coalesce(sqlc.narg('location'), location),
  password = coalesce(sqlc.narg('password'), password),
  updated_at = NOW()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteCompany :exec
DELETE FROM companies
WHERE id = $1;
