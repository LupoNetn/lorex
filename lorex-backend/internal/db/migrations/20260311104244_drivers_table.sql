-- +goose Up
CREATE TABLE IF NOT EXISTS drivers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone TEXT NOT NULL,
    password TEXT NOT NULL,
    dob DATE NOT NULL,
    gender TEXT NOT NULL,
    state_residence TEXT NOT NULL,
    country_residence TEXT NOT NULL,
    nationality TEXT NOT NULL,
    available BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_drivers_available ON drivers(available);
CREATE INDEX IF NOT EXISTS idx_drivers_residence ON drivers(country_residence, state_residence);

-- +goose Down
DROP INDEX IF EXISTS idx_drivers_available;
DROP INDEX IF EXISTS idx_drivers_residence;
DROP TABLE IF EXISTS drivers;
