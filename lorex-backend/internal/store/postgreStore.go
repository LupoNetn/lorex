package store

import (
	"context"

	"github.com/luponetn/lorex/internal/auth"
	"github.com/luponetn/lorex/internal/db/sqlc"
)

type PostgresStore struct {
	db sqlc.Querier
}

func NewPostgresStore(db sqlc.Querier) *PostgresStore {
	return &PostgresStore{
		db: db,
	}
}
func (s *PostgresStore) CreateCompany(ctx context.Context, arg sqlc.CreateCompanyParams) (sqlc.Company, error) {
	return s.db.CreateCompany(ctx, arg)
}

func (s *PostgresStore) LoginCompany(ctx context.Context, arg auth.LoginCompanyRequest) (auth.LoginCompanyResponse, error) {
	// Implement login logic here or proxy it to sqlc
	return auth.LoginCompanyResponse{}, nil
}


