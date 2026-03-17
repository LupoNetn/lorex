package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/luponetn/lorex/internal/db/sqlc"
	"github.com/luponetn/lorex/internal/location"
)

// TaskStore defines the minimal DB methods needed for background jobs.
// This makes mocking for testing extremely easy.
type TaskStore interface {
	GetDelivery(ctx context.Context, id pgtype.UUID) (sqlc.Delivery, error)
	UpdateDeliveryStatus(ctx context.Context, arg sqlc.UpdateDeliveryStatusParams) (sqlc.Delivery, error)
	GetDriver(ctx context.Context, id pgtype.UUID) (sqlc.Driver, error)
}

type TaskProcessor interface {
	Start() error
	ProcessTaskAssignDriver(ctx context.Context, t *asynq.Task) error
	AssignDriver(ctx context.Context, deliveryID string) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  TaskStore
	locStore location.Store
}

func NewRedisTaskProcessor(redisOpt asynq.RedisConnOpt, store TaskStore, locStore location.Store) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Concurrency: 10,
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				slog.Error("task processing failed", "type", task.Type(), "error", err)
			}),
		},
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
		locStore: locStore,
	}
}

func (p *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	// Route tasks to their handlers
	mux.HandleFunc(TypeAssignDriver, p.ProcessTaskAssignDriver)

	return p.server.Start(mux)
}

func (p *RedisTaskProcessor) ProcessTaskAssignDriver(ctx context.Context, t *asynq.Task) error {
	var payload AssignDriverPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	slog.Info("processing driver assignment", "delivery_id", payload.DeliveryID)

	return p.AssignDriver(ctx, payload.DeliveryID)
}

func (p *RedisTaskProcessor) AssignDriver(ctx context.Context, deliveryID string) error {
	//get details about full delivery
	var uuid pgtype.UUID
	uuid.Scan(deliveryID)
	delivery, err := p.store.GetDelivery(ctx, uuid)
	if err != nil {
		slog.Error("could not get delivery", "error", err)
		return err
	}

	//get nearest driver to delivery pickup point
	drivers, err := p.locStore.GetNearbyDrivers(ctx, delivery.PickupLat.Float64, delivery.PickupLng.Float64, 5.0)
	if err != nil {
		slog.Error("could not get nearby drivers", "error", err)
		return err
	}

	if len(drivers) == 0 {
		slog.Warn("no drivers found for delivery", "delivery_id", deliveryID)
		return nil
	}
    
	var driverScore int
	var computedDriverScore int
	var assignedDriverID string
	for _, driver := range drivers {
		var driverUUID pgtype.UUID
		driverUUID.Scan(driver)
        fetchedDriver, err := p.store.GetDriver(ctx,driverUUID)
		if err != nil {
			slog.Error("could not get driver", "error", err)
			continue
		}
		if fetchedDriver.ActiveDeliveryID.Valid {
			continue
		}
		if fetchedDriver.VehicleType != delivery.SuitableVehicle {
			continue
		}
		if fetchedDriver.MaxWeightCapacity < delivery.Weight.Float64 {
			continue
		}
		computedDriverScore = 1
		if fetchedDriver.Rating > 4 {
			computedDriverScore *= 10
		}
		if fetchedDriver.TotalDeliveries > 10 {
			computedDriverScore++
		}
		if computedDriverScore > driverScore {
			driverScore = computedDriverScore
			assignedDriverID = driverUUID.String()
		}
		//fix
	}
	if assignedDriverID == "" {
		slog.Warn("no drivers found for delivery", "delivery_id", deliveryID)
		return nil
	}
	

	return nil
}
