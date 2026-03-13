package delivery

import (
	"context"
	"encoding/hex"
	"log/slog"

	"github.com/luponetn/lorex/internal/db/sqlc"
	"github.com/luponetn/lorex/internal/tasks"
)

// type DeliveryStore interface {
// 	CreateDelivery(ctx context.Context, arg sqlc.CreateDeliveryParams) (sqlc.Delivery, error)
// }

type Service interface {
	CreateDelivery(ctx context.Context, arg sqlc.CreateDeliveryParams) (sqlc.Delivery, error)
}

type Svc struct {
	store  Service
	enquer *tasks.AsynqClient
}

func NewSvc(store Service, enquer *tasks.AsynqClient) Service {
	return &Svc{
		store:  store,
		enquer: enquer,
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
