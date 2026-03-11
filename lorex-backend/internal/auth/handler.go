package auth

import (
	"github.com/gin-gonic/gin"
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
func (h *Handler) CompanySignup(c *gin.Context) {}

func (h *Handler) CompanyLogin(c *gin.Context) {}