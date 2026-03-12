package auth

import (
	"time"
	"github.com/google/uuid"
)

type CreateCompanyRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required email"`
	Password    string `json:"password" binding:"required min=7"`
	Phone       string `json:"phone" binding:"required min=11 max=15"`
	Industry    string `json:"industry" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Description string `json:"description" binding:"required"`
	Plan        string `json:"plan" binding:"required"`
}

type LoginCompanyRequest struct {
	Email    string `json:"email" binding:"required email"`
	Password string `json:"password" binding:"required min=7"`
}

type CompanyResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Industry  string    `json:"industry"`
	Location  string    `json:"location"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginCompanyResponse struct {
	AccessToken  string          `json:"access_token"`
	RefreshToken string          `json:"refresh_token"`
	Company      CompanyResponse `json:"company"`
}

type CreateDriverRequest struct {
	CompanyCode      string `json:"company_code" binding:"required"`
	Name             string `json:"name" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	Phone            string `json:"phone" binding:"required,min=11,max=15"`
	Password         string `json:"password" binding:"required,min=7"`
	Dob              string `json:"dob" binding:"required"` // Assuming string YYYY-MM-DD for simpler binding
	Gender           string `json:"gender" binding:"required"`
	StateResidence   string `json:"state_residence" binding:"required"`
	CountryResidence string `json:"country_residence" binding:"required"`
	Nationality      string `json:"nationality" binding:"required"`
}

type LoginDriverRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=7"`
}

type DriverResponse struct {
	ID               uuid.UUID `json:"id"`
	CompanyID        uuid.UUID `json:"company_id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	Dob              string    `json:"dob"`
	Gender           string    `json:"gender"`
	StateResidence   string    `json:"state_residence"`
	CountryResidence string    `json:"country_residence"`
	Nationality      string    `json:"nationality"`
	Available        bool      `json:"available"`
	CreatedAt        time.Time `json:"created_at"`
}

type LoginDriverResponse struct {
	AccessToken  string         `json:"access_token"`
	RefreshToken string         `json:"refresh_token"`
	Driver       DriverResponse `json:"driver"`
}

type CreateCustomerRequest struct {
	CompanyCode      string `json:"company_code" binding:"required"`
	Name             string `json:"name" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	Phone            string `json:"phone" binding:"required,min=11,max=15"`
	Password         string `json:"password" binding:"required,min=7"`
	Dob              string `json:"dob" binding:"required"` 
	Gender           string `json:"gender" binding:"required"`
	StateResidence   string `json:"state_residence" binding:"required"`
	CountryResidence string `json:"country_residence" binding:"required"`
	Nationality      string `json:"nationality" binding:"required"`
}

type LoginCustomerRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=7"`
}

type CustomerResponse struct {
	ID               uuid.UUID `json:"id"`
	CompanyID        uuid.UUID `json:"company_id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	Dob              string    `json:"dob"`
	Gender           string    `json:"gender"`
	StateResidence   string    `json:"state_residence"`
	CountryResidence string    `json:"country_residence"`
	Nationality      string    `json:"nationality"`
	CreatedAt        time.Time `json:"created_at"`
}

type LoginCustomerResponse struct {
	AccessToken  string           `json:"access_token"`
	RefreshToken string           `json:"refresh_token"`
	Customer     CustomerResponse `json:"customer"`
}