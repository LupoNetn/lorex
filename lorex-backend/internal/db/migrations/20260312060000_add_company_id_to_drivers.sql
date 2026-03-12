-- +goose Up
ALTER TABLE drivers ADD COLUMN company_id UUID REFERENCES companies(id);
CREATE INDEX IF NOT EXISTS idx_drivers_company_id ON drivers(company_id);

-- +goose Down
DROP INDEX IF EXISTS idx_drivers_company_id;
ALTER TABLE drivers DROP COLUMN IF EXISTS company_id;
