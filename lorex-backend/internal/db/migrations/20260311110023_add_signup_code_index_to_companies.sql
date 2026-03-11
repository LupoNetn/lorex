-- +goose Up
-- Add a unique index for customer_signup_code in the companies table
CREATE UNIQUE INDEX IF NOT EXISTS idx_companies_customer_signup_code ON companies(customer_signup_code);

-- +goose Down
-- Remove the index if we roll back
DROP INDEX IF EXISTS idx_companies_customer_signup_code;
