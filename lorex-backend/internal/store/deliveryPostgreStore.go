package store

import (
	"context"

	"github.com/luponetn/lorex/internal/db/sqlc"
)

type DeliveryPostgresStore struct {
	db sqlc.Querier
}

func NewDeliveryPostgresStore(db sqlc.Querier) *DeliveryPostgresStore {
	return &DeliveryPostgresStore{
		db: db,
	}
}

func (s *DeliveryPostgresStore) CreateDelivery(ctx context.Context, arg sqlc.CreateDeliveryParams) (sqlc.Delivery, error) {
	return s.db.CreateDelivery(ctx, arg)
}