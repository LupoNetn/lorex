package notifications

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/luponetn/lorex/internal/db/sqlc"
)

type NotificationStore interface {
	CreateNotification(ctx context.Context, arg sqlc.CreateNotificationParams) error
	GetNotifications(ctx context.Context, receiverID pgtype.UUID) ([]sqlc.Notification, error)
	MarkNotificationAsRead(ctx context.Context, id pgtype.UUID) error
}

type Service interface {
	CreateNotification(ctx context.Context, arg sqlc.CreateNotificationParams) error
	Send(ctx context.Context, receiverID pgtype.UUID, receiverType string, companyID pgtype.UUID, message string, nType string) error
	GetNotifications(ctx context.Context, receiverID pgtype.UUID) ([]sqlc.Notification, error)
	MarkAsRead(ctx context.Context, id pgtype.UUID) error
}

type Svc struct {
	store NotificationStore
}

func NewSvc(store NotificationStore) Service {
	return &Svc{
		store: store,
	}
}

func (s *Svc) CreateNotification(ctx context.Context, arg sqlc.CreateNotificationParams) error {
	return s.store.CreateNotification(ctx, arg)
}

func (s *Svc) Send(ctx context.Context, receiverID pgtype.UUID, receiverType string, companyID pgtype.UUID, message string, nType string) error {
	return s.store.CreateNotification(ctx, sqlc.CreateNotificationParams{
		ReceiverID:   receiverID,
		ReceiverType: receiverType,
		CompanyID:    companyID,
		Message:      message,
		Type:         nType,
		Read:         pgtype.Bool{Bool: false, Valid: true},
	})
}

func (s *Svc) GetNotifications(ctx context.Context, receiverID pgtype.UUID) ([]sqlc.Notification, error) {
	return s.store.GetNotifications(ctx, receiverID)
}

func (s *Svc) MarkAsRead(ctx context.Context, id pgtype.UUID) error {
	return s.store.MarkNotificationAsRead(ctx, id)
}
