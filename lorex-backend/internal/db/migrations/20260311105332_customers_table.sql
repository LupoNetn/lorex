-- +goose Up
CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    company_id UUID REFERENCES companies(id),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone TEXT NOT NULL,
    password TEXT NOT NULL,
    dob DATE NOT NULL,
    gender TEXT NOT NULL,
    state_residence TEXT NOT NULL,
    country_residence TEXT NOT NULL,
    nationality TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_customers_company_id ON customers(company_id);
CREATE INDEX IF NOT EXISTS idx_customers_residence ON customers(country_residence, state_residence);

-- +goose Down
DROP INDEX IF EXISTS idx_customers_company_id;
DROP INDEX IF EXISTS idx_customers_residence;
DROP TABLE IF EXISTS customers;
