package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/luponetn/lorex/internal/db/sqlc"
)

type NotificationPostgresStore struct {
	db sqlc.Querier
}

func NewNotificationPostgresStore(db sqlc.Querier) *NotificationPostgresStore {
	return &NotificationPostgresStore{
		db: db,
	}
}

func (s *NotificationPostgresStore) CreateNotification(ctx context.Context, arg sqlc.CreateNotificationParams) error {
	return s.db.CreateNotification(ctx, arg)
}

func (s *NotificationPostgresStore) GetNotifications(ctx context.Context, receiverID pgtype.UUID) ([]sqlc.Notification, error) {
	return s.db.GetNotifications(ctx, receiverID)
}

func (s *NotificationPostgresStore) MarkNotificationAsRead(ctx context.Context, id pgtype.UUID) error {
	return s.db.MarkNotificationAsRead(ctx, id)
}
