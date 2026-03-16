-- +goose Up
-- +goose StatementBegin
DO $$ BEGIN
  CREATE TYPE status AS ENUM ('pending', 'assigned', 'successful', 'failed');
EXCEPTION 
  WHEN duplicate_object THEN null;
END $$; 
-- +goose StatementEnd 

CREATE TABLE IF NOT EXISTS deliveries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id UUID NOT NULL REFERENCES customers(id),
    driver_id UUID REFERENCES drivers(id),
    pickup_address TEXT NOT NULL,
    delivery_address TEXT NOT NULL,
    pickup_lat DOUBLE PRECISION,
    pickup_lng DOUBLE PRECISION,
    delivery_lat DOUBLE PRECISION,
    delivery_lng DOUBLE PRECISION,
    status TEXT NOT NULL DEFAULT 'pending',
    package_type TEXT,
    weight DOUBLE PRECISION,
    price DECIMAL(10, 2),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_deliveries_customer_id ON deliveries(customer_id);
CREATE INDEX IF NOT EXISTS idx_deliveries_driver_id ON deliveries(driver_id);
CREATE INDEX IF NOT EXISTS idx_deliveries_status ON deliveries(status);

-- +goose Down
DROP INDEX IF EXISTS idx_deliveries_status;
DROP INDEX IF EXISTS idx_deliveries_driver_id;
DROP INDEX IF EXISTS idx_deliveries_customer_id;
DROP TABLE IF EXISTS deliveries;
