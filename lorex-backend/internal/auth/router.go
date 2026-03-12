package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
   companyAuthGroup := r.Group("/company-auth")
   driverAuthGroup := r.Group("/driver-auth")
   customerAuthGroup := r.Group("/customer-auth")

   //company auth routes
   companyAuthGroup.POST("/signup", h.CompanySignup)
   companyAuthGroup.POST("/login", h.CompanyLogin)

   //driver auth routes
   driverAuthGroup.POST("/signup", h.DriverSignup)
   driverAuthGroup.POST("/login", h.DriverLogin)

   //customer auth routes
   customerAuthGroup.POST("/signup", h.CustomerSignup)
   customerAuthGroup.POST("/login", h.CustomerLogin)
}