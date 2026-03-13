package tasks

import (
	"encoding/json"
	"log/slog"

	"github.com/hibiken/asynq"
)

// A list of background tasks types
const (
	TypeAssignDriver = "driver:assign"
)

// functions to create new background tasks.
func NewAssignDriverTask(deliveryID string) (*asynq.Task, error) {
	payload, err := json.Marshal(AssignDriverPayload{
		DeliveryID: deliveryID,
	})
	if err != nil {
		slog.Error("unable to create new task for assigning driver", "error", err)
	}

	return asynq.NewTask(TypeAssignDriver, payload), nil
}
