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

func (s *PostgresStore) GetCompanyBySignupCode(ctx context.Context, signupCode string) (sqlc.Company, error) {
	return s.db.GetCompanyBySignupCode(ctx, signupCode)
}

func (s *PostgresStore) CreateDriver(ctx context.Context, arg sqlc.CreateDriverParams) (sqlc.Driver, error) {
	return s.db.CreateDriver(ctx, arg)
}

func (s *PostgresStore) GetDriverByEmail(ctx context.Context, email string) (sqlc.Driver, error) {
	return s.db.GetDriverByEmail(ctx, email)
}

func (s *PostgresStore) LoginDriver(ctx context.Context, arg auth.LoginDriverRequest) (auth.LoginDriverResponse, error) {
	return auth.LoginDriverResponse{}, errors.New("not implemented in store")
}

func (s *PostgresStore) CreateCustomer(ctx context.Context, arg sqlc.CreateCustomerParams) (sqlc.Customer, error) {
	return s.db.CreateCustomer(ctx, arg)
}

func (s *PostgresStore) GetCustomerByEmail(ctx context.Context, email string) (sqlc.Customer, error) {
	return s.db.GetCustomerByEmail(ctx, email)
}

func (s *PostgresStore) LoginCustomer(ctx context.Context, arg auth.LoginCustomerRequest) (auth.LoginCustomerResponse, error) {
	return auth.LoginCustomerResponse{}, errors.New("not implemented in store")
}
