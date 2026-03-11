package auth

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/luponetn/lorex/internal/db/sqlc"
	"github.com/luponetn/lorex/utils"
)

type Handler struct {
	Svc Service
}

func NewHandler(svc Service) *Handler {
  return &Handler{
	Svc: svc,
  }
}

// implement handlers for auth group
func (h *Handler) CompanySignup(c *gin.Context) {
	var req CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Could not properly bind request with json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationError(err)})
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		slog.Error("Could not hash password", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Prepare the arguments for creating a company
	arg := sqlc.CreateCompanyParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Phone:    req.Phone,
		Industry: req.Industry,
		Location: req.Location,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  true,
		},
		// For now, let's just generate a simple signup code or leave it placeholder
		CustomerSignupCode: "CO-" + req.Name[:3],
	}

	// Call the service to create the company
	company, err := h.Svc.CreateCompany(c.Request.Context(), arg)
	if err != nil {
		slog.Error("Could not create company", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
		return
	}

	// Respond with the created company (avoiding returning the password)
	res := CompanyResponse{
		ID:        company.ID.Bytes, // uuid.UUID is likely compatible or needs conversion from pgtype.UUID
		Name:      company.Name,
		Email:     company.Email,
		Industry:  company.Industry,
		Location:  company.Location,
		Phone:     company.Phone,
		CreatedAt: company.CreatedAt.Time,
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) CompanyLogin(c *gin.Context) {
	var req LoginCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationError(err)})
		return
	}

	// Call service to get login details (usually including tokens and company info)
	res, err := h.Svc.LoginCompany(c.Request.Context(), req)
	if err != nil {
		slog.Error("Login failed", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, res)
}