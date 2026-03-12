package auth

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/luponetn/lorex/internal/db/sqlc"
	"github.com/luponetn/lorex/utils"
	"time"
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
		ID:        company.ID.Bytes, 
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

func (h *Handler) DriverSignup(c *gin.Context) {
	var req CreateDriverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Could not properly bind request with json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationError(err)})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		slog.Error("Could not hash password", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	dobTime, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DOB format, expected YYYY-MM-DD"})
		return
	}

	company, err := h.Svc.GetCompanyBySignupCode(c.Request.Context(), req.CompanyCode)
	if err != nil {
		slog.Error("Invalid company code", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company code"})
		return
	}

	arg := sqlc.CreateDriverParams{
		CompanyID:        pgtype.UUID{Bytes: company.ID.Bytes, Valid: true},
		Name:             req.Name,
		Email:            req.Email,
		Phone:            req.Phone,
		Password:         hashedPassword,
		Dob:              pgtype.Date{Time: dobTime, Valid: true},
		Gender:           req.Gender,
		StateResidence:   req.StateResidence,
		CountryResidence: req.CountryResidence,
		Nationality:      req.Nationality,
	}

	driver, err := h.Svc.CreateDriver(c.Request.Context(), arg)
	if err != nil {
		slog.Error("Could not create driver", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create driver"})
		return
	}

	res := DriverResponse{
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
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) DriverLogin(c *gin.Context) {
	var req LoginDriverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationError(err)})
		return
	}

	res, err := h.Svc.LoginDriver(c.Request.Context(), req)
	if err != nil {
		slog.Error("Login failed", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) CustomerSignup(c *gin.Context) {
	var req CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Could not properly bind request with json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationError(err)})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		slog.Error("Could not hash password", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	dobTime, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DOB format, expected YYYY-MM-DD"})
		return
	}

	company, err := h.Svc.GetCompanyBySignupCode(c.Request.Context(), req.CompanyCode)
	if err != nil {
		slog.Error("Invalid company code", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company code"})
		return
	}

	arg := sqlc.CreateCustomerParams{
		CompanyID:        pgtype.UUID{Bytes: company.ID.Bytes, Valid: true},
		Name:             req.Name,
		Email:            req.Email,
		Phone:            req.Phone,
		Password:         hashedPassword,
		Dob:              pgtype.Date{Time: dobTime, Valid: true},
		Gender:           req.Gender,
		StateResidence:   req.StateResidence,
		CountryResidence: req.CountryResidence,
		Nationality:      req.Nationality,
	}

	customer, err := h.Svc.CreateCustomer(c.Request.Context(), arg)
	if err != nil {
		slog.Error("Could not create customer", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	res := CustomerResponse{
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
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) CustomerLogin(c *gin.Context) {
	var req LoginCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationError(err)})
		return
	}

	res, err := h.Svc.LoginCustomer(c.Request.Context(), req)
	if err != nil {
		slog.Error("Login failed", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, res)
}