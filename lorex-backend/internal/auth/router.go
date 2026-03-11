package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
   authGroup := r.Group("/auth")

   authGroup.POST("/company/signup", h.CompanySignup)
   authGroup.POST("/company/login", h.CompanyLogin)
}