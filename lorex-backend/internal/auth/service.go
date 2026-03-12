package auth

import (
	"context"
	"errors"

	"github.com/luponetn/lorex/internal/db/sqlc"
	"github.com/luponetn/lorex/utils"
	"time"
)

type Service interface {
	CreateCompany(ctx context.Context, arg sqlc.CreateCompanyParams) (sqlc.Company, error)
	LoginCompany(ctx context.Context, arg LoginCompanyRequest) (LoginCompanyResponse, error)
	GetCompanyByEmail(ctx context.Context, email string) (sqlc.Company, error)
	GetCompanyBySignupCode(ctx context.Context, signupCode string) (sqlc.Company, error)

	CreateDriver(ctx context.Context, arg sqlc.CreateDriverParams) (sqlc.Driver, error)
	LoginDriver(ctx context.Context, arg LoginDriverRequest) (LoginDriverResponse, error)
	GetDriverByEmail(ctx context.Context, email string) (sqlc.Driver, error)

	CreateCustomer(ctx context.Context, arg sqlc.CreateCustomerParams) (sqlc.Customer, error)
	LoginCustomer(ctx context.Context, arg LoginCustomerRequest) (LoginCustomerResponse, error)
	GetCustomerByEmail(ctx context.Context, email string) (sqlc.Customer, error)
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

/* TODO: Add Email verification and background jobs processing later on!!*/
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
	accessToken, refreshToken, err := utils.GenerateTokens(company.ID.Bytes, company.Email, company.ID.Bytes)
	
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

func (s *Svc) GetCompanyBySignupCode(ctx context.Context, signupCode string) (sqlc.Company, error) {
	return s.store.GetCompanyBySignupCode(ctx, signupCode)
}

func (s *Svc) CreateDriver(ctx context.Context, arg sqlc.CreateDriverParams) (sqlc.Driver, error) {
   return s.store.CreateDriver(ctx,arg)
}

func (s *Svc) LoginDriver(ctx context.Context, arg LoginDriverRequest) (LoginDriverResponse, error) {
	driver, err := s.store.GetDriverByEmail(ctx, arg.Email)
	if err != nil {
		return LoginDriverResponse{}, errors.New("invalid credentials")
	}

	err = utils.ComparePassword(arg.Password, driver.Password)
	if err != nil {
		return LoginDriverResponse{}, errors.New("invalid credentials")
	}

	// Generate JWT tokens
	accessToken, refreshToken, err := utils.GenerateTokens(driver.ID.Bytes, driver.Email, driver.ID.Bytes)
	
	if err != nil {
		return LoginDriverResponse{}, errors.New("failed to generate auth tokens")
	}

	return LoginDriverResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Driver: DriverResponse{
			ID:               driver.ID.Bytes,
			CompanyID:        driver.CompanyID.Bytes,
			Name:             driver.Name,
			Email:            driver.Email,
			Phone:            driver.Phone,
			Dob:              driver.Dob.Time.Format(time.DateOnly),
			Gender:           driver.Gender,
			StateResidence:   driver.StateResidence,
			CountryResidence: driver.CountryResidence,
			Nationality:      driver.Nationality,
			Available:        driver.Available.Bool,
			CreatedAt:        driver.CreatedAt.Time,
		},
	}, nil
}

func (s *Svc) GetDriverByEmail(ctx context.Context, email string) (sqlc.Driver, error) {
	return s.store.GetDriverByEmail(ctx, email)
}

func (s *Svc) CreateCustomer(ctx context.Context, arg sqlc.CreateCustomerParams) (sqlc.Customer, error) {
   return s.store.CreateCustomer(ctx,arg)
}

func (s *Svc) LoginCustomer(ctx context.Context, arg LoginCustomerRequest) (LoginCustomerResponse, error) {
	customer, err := s.store.GetCustomerByEmail(ctx, arg.Email)
	if err != nil {
		return LoginCustomerResponse{}, errors.New("invalid credentials")
	}

	err = utils.ComparePassword(arg.Password, customer.Password)
	if err != nil {
		return LoginCustomerResponse{}, errors.New("invalid credentials")
	}

	accessToken, refreshToken, err := utils.GenerateTokens(customer.ID.Bytes, customer.Email, customer.ID.Bytes)
	
	if err != nil {
		return LoginCustomerResponse{}, errors.New("failed to generate auth tokens")
	}

	return LoginCustomerResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Customer: CustomerResponse{
			ID:               customer.ID.Bytes,
			CompanyID:        customer.CompanyID.Bytes,
			Name:             customer.Name,
			Email:            customer.Email,
			Phone:            customer.Phone,
			Dob:              customer.Dob.Time.Format(time.DateOnly),
			Gender:           customer.Gender,
			StateResidence:   customer.StateResidence,
			CountryResidence: customer.CountryResidence,
			Nationality:      customer.Nationality,
			CreatedAt:        customer.CreatedAt.Time,
		},
	}, nil
}

func (s *Svc) GetCustomerByEmail(ctx context.Context, email string) (sqlc.Customer, error) {
	return s.store.GetCustomerByEmail(ctx, email)
}