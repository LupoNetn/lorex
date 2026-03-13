package store

import (
	"context"
	"errors"

	"github.com/luponetn/lorex/internal/db/sqlc"
	"github.com/luponetn/lorex/internal/auth"
)

type AuthPostgresStore struct {
	db sqlc.Querier
}

func NewAuthPostgresStore(db sqlc.Querier) *AuthPostgresStore {
	return &AuthPostgresStore{
		db: db,
	}
}
func (s *AuthPostgresStore) CreateCompany(ctx context.Context, arg sqlc.CreateCompanyParams) (sqlc.Company, error) {
	return s.db.CreateCompany(ctx, arg)
}

func (s *AuthPostgresStore) GetCompanyByEmail(ctx context.Context, email string) (sqlc.Company, error) {
	return s.db.GetCompanyByEmail(ctx, email)
}

func (s *AuthPostgresStore) LoginCompany(ctx context.Context, arg auth.LoginCompanyRequest) (auth.LoginCompanyResponse, error) {
	return auth.LoginCompanyResponse{}, errors.New("not implemented in store")
}

func (s *AuthPostgresStore) GetCompanyBySignupCode(ctx context.Context, signupCode string) (sqlc.Company, error) {
	return s.db.GetCompanyBySignupCode(ctx, signupCode)
}

func (s *AuthPostgresStore) CreateDriver(ctx context.Context, arg sqlc.CreateDriverParams) (sqlc.Driver, error) {
	return s.db.CreateDriver(ctx, arg)
}

func (s *AuthPostgresStore) GetDriverByEmail(ctx context.Context, email string) (sqlc.Driver, error) {
	return s.db.GetDriverByEmail(ctx, email)
}

func (s *AuthPostgresStore) LoginDriver(ctx context.Context, arg auth.LoginDriverRequest) (auth.LoginDriverResponse, error) {
	return auth.LoginDriverResponse{}, errors.New("not implemented in store")
}

func (s *AuthPostgresStore) CreateCustomer(ctx context.Context, arg sqlc.CreateCustomerParams) (sqlc.Customer, error) {
	return s.db.CreateCustomer(ctx, arg)
}

func (s *AuthPostgresStore) GetCustomerByEmail(ctx context.Context, email string) (sqlc.Customer, error) {
	return s.db.GetCustomerByEmail(ctx, email)
}

func (s *AuthPostgresStore) LoginCustomer(ctx context.Context, arg auth.LoginCustomerRequest) (auth.LoginCustomerResponse, error) {
	return auth.LoginCustomerResponse{}, errors.New("not implemented in store")
}
