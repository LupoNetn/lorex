package delivery

import (
	"context"
	"encoding/hex"
	"log/slog"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/luponetn/lorex/internal/db/sqlc"
	"github.com/luponetn/lorex/internal/location"
	"github.com/luponetn/lorex/internal/tasks"
)

type DeliveryStore interface {
	CreateDelivery(ctx context.Context, arg sqlc.CreateDeliveryParams) (sqlc.Delivery, error)
	AssignDriver(ctx context.Context, arg sqlc.AssignDriverParams) (sqlc.Delivery, error)
	GetDelivery(ctx context.Context, deliveryID pgtype.UUID) (sqlc.Delivery, error)
}

type Service interface {
	CreateDelivery(ctx context.Context, arg sqlc.CreateDeliveryParams) (sqlc.Delivery, error)
}

type Svc struct {
	store  DeliveryStore
	enquer *tasks.AsynqClient
	locStore location.Store
}

func NewSvc(store DeliveryStore, enquer *tasks.AsynqClient, locStore location.Store) Service {
	return &Svc{
		store:  store,
		enquer: enquer,
		locStore: locStore,
	}
}

// implement delivery functionality
func (s *Svc) CreateDelivery(ctx context.Context, arg sqlc.CreateDeliveryParams) (sqlc.Delivery, error) {
	delivery, err := s.store.CreateDelivery(ctx, arg)
	if err != nil {
		slog.Error("could not create delivery", "error", err)
		return sqlc.Delivery{}, err
	}

	// CRITICAL: Push to background job
	s.enquer.DistributeAssignDriverTask(hex.EncodeToString(delivery.ID.Bytes[:]))

	slog.Info("delivery created and assignment task enqueued", "delivery_id", hex.EncodeToString(delivery.ID.Bytes[:]))
	return delivery, nil
}

