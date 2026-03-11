package auth

import (
	"context"
	"errors"

	"github.com/luponetn/lorex/internal/db/sqlc"
	"github.com/luponetn/lorex/utils"
)

type Service interface {
	CreateCompany(ctx context.Context, arg sqlc.CreateCompanyParams) (sqlc.Company, error)
	LoginCompany(ctx context.Context, arg LoginCompanyRequest) (LoginCompanyResponse, error)
	GetCompanyByEmail(ctx context.Context, email string) (sqlc.Company, error)
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
   return s.store.CreateCompany(ctx,arg)
}

func (s *Svc) LoginCompany(ctx context.Context, arg LoginCompanyRequest) (LoginCompanyResponse, error) {
	company, err := s.store.GetCompanyByEmail(ctx, arg.Email)
	if err != nil {
		return LoginCompanyResponse{}, errors.New("invalid credentials")
	}

	err = utils.ComparePassword(arg.Password, company.Password)
	if err != nil {
		return LoginCompanyResponse{}, errors.New("invalid credentials")
	}

	// Generate JWT tokens
	accessToken, refreshToken, err := utils.GenerateTokens(company.ID.Bytes, company.Email)
	if err != nil {
		return LoginCompanyResponse{}, errors.New("failed to generate auth tokens")
	}

	return LoginCompanyResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Company: CompanyResponse{
			ID:        company.ID.Bytes,
			Name:      company.Name,
			Email:     company.Email,
			Industry:  company.Industry,
			Location:  company.Location,
			Phone:     company.Phone,
			CreatedAt: company.CreatedAt.Time,
		},
	}, nil
}

func (s *Svc) GetCompanyByEmail(ctx context.Context, email string) (sqlc.Company, error) {
	return s.store.GetCompanyByEmail(ctx, email)
}