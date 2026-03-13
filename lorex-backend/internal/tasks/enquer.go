package tasks

import (
	"log/slog"

	"github.com/hibiken/asynq"
)

type AsynqClient struct {
	client *asynq.Client
}

func NewAsynqClient(client *asynq.Client) *AsynqClient {
	return &AsynqClient{
		client: client,
	}
}

//distribute tasks
func (a *AsynqClient) DistributeAssignDriverTask(deliveryID string) {
	task, err := NewAssignDriverTask(deliveryID)
	if err != nil {
       slog.Error("an error occured when trying to create assign driver task for distribution", "error", err)
	   return
	}

	info, err := a.client.Enqueue(task, asynq.MaxRetry(4))
	if err != nil {
		slog.Error("an error occured when trying to enque assign driver task")
	}

	slog.Info("Enqued tasks details", "id", info.ID, "queue", info.Queue, "type", info.Type)
}