-- name: GetCompany :one
SELECT * FROM companies
WHERE id = $1 LIMIT 1;

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
  set name = $2,
  description = $3,
  logo = $4,
  industry = $5,
  phone = $6,
  email = $7,
  location = $8,
  password = $9,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCompany :exec
DELETE FROM companies
WHERE id = $1;
