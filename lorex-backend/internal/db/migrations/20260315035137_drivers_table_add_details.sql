-- +goose Up
CREATE TYPE vehicle_enum AS ENUM ('bike', 'motorcycle', 'car', 'van', 'truck');

ALTER TABLE drivers
ADD COLUMN active_delivery_id UUID REFERENCES deliveries(id) ON DELETE SET NULL,
ADD COLUMN vehicle_type vehicle_enum NOT NULL DEFAULT 'bike', 
ADD COLUMN vehicle_plate_number TEXT NOT NULL DEFAULT 'ABC-123',
ADD COLUMN max_weight_capacity FLOAT NOT NULL DEFAULT 100.0,
ADD COLUMN rating FLOAT NOT NULL DEFAULT 0.0,
ADD COLUMN total_deliveries INT NOT NULL DEFAULT 0;

CREATE INDEX idx_drivers_vehicle_type ON drivers(vehicle_type);

-- +goose Down
ALTER TABLE drivers
DROP COLUMN active_delivery_id,
DROP COLUMN vehicle_type,
DROP COLUMN vehicle_plate_number,
DROP COLUMN max_weight_capacity,
DROP COLUMN rating,
DROP COLUMN total_deliveries;

DROP INDEX idx_drivers_vehicle_type;
