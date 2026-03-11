package auth

import (
	"context"

	"github.com/luponetn/lorex/internal/db/sqlc"
)

type Service interface {
	CreateCompany(ctx context.Context, arg sqlc.CreateCompanyParams) (sqlc.Company, error)
	LoginCompany(ctx context.Context, arg LoginCompanyRequest) (LoginCompanyResponse, error)
}

type Svc struct {
	store Service
}

//construct Svc struct
func NewSvc(store Service) Service {
	return &Svc{
		store: store,
	}
}

//functions for svc struct to implement service 
func (s *Svc) CreateCompany(ctx context.Context, arg sqlc.CreateCompanyParams) (sqlc.Company, error) {
	return sqlc.Company{}, nil
}

func (s *Svc) LoginCompany(ctx context.Context, arg LoginCompanyRequest) (LoginCompanyResponse, error) {
	return LoginCompanyResponse{}, nil
}