package location

import (
	"github.com/gin-gonic/gin"
	"github.com/luponetn/lorex/internal/middleware"
)

func RegisterRoutes(router *gin.Engine, h *Handler) {
	locationGroup := router.Group("/location")
	locationGroup.Use(middleware.AuthMiddleware())
	locationGroup.Use(middleware.DriverOnly())

	locationGroup.POST("/heartbeat", h.UpdateLocation)
}
