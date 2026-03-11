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