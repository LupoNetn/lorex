package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/luponetn/lorex/internal/db/sqlc"
)

// TaskStore defines the minimal DB methods needed for background jobs.
// This makes mocking for testing extremely easy.
type TaskStore interface {
	GetDelivery(ctx context.Context, id pgtype.UUID) (sqlc.Delivery, error)
	UpdateDeliveryStatus(ctx context.Context, arg sqlc.UpdateDeliveryStatusParams) (sqlc.Delivery, error)
}

type TaskProcessor interface {
	Start() error
	ProcessTaskAssignDriver(ctx context.Context, t *asynq.Task) error
	AssignDriver(ctx context.Context, deliveryID string) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  TaskStore
}

func NewRedisTaskProcessor(redisOpt asynq.RedisConnOpt, store TaskStore) TaskProcessor {
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
	// Logic for finding and assigning a driver will go here
	// This method only takes the deliveryID, as requested.
	return nil
}
