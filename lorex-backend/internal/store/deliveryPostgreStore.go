package store

import (
	"context"

	"github.com/luponetn/lorex/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
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

func (s *DeliveryPostgresStore) AssignDriver(ctx context.Context, arg sqlc.AssignDriverParams) (sqlc.Delivery, error) {
	return s.db.AssignDriver(ctx, arg)
}

func (s *DeliveryPostgresStore) GetDelivery(ctx context.Context, deliveryID pgtype.UUID) (sqlc.Delivery, error) {
	return s.db.GetDelivery(ctx,deliveryID)
}