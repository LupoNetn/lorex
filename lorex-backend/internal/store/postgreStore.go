package store

import (
	"context"
	"errors"

	"github.com/luponetn/lorex/internal/db/sqlc"
	"github.com/luponetn/lorex/internal/auth"
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

func (s *PostgresStore) GetCompanyByEmail(ctx context.Context, email string) (sqlc.Company, error) {
	return s.db.GetCompanyByEmail(ctx, email)
}

func (s *PostgresStore) LoginCompany(ctx context.Context, arg auth.LoginCompanyRequest) (auth.LoginCompanyResponse, error) {
	return auth.LoginCompanyResponse{}, errors.New("not implemented in store")
}


