-- +goose Up
ALTER TABLE deliveries
ADD COLUMN suitable_vehicle vehicle_enum NOT NULL DEFAULT 'bike';

-- +goose Down
ALTER TABLE deliveries
DROP COLUMN suitable_vehicle;
