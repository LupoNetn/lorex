package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/luponetn/lorex/internal/middleware"
)


func RegisterRoutes(router *gin.Engine, h *Handler) {
	deliveryGroup := router.Group("/delivery")
	deliveryGroup.Use(middleware.AuthMiddleware())

	//implement delivery routes
	deliveryGroup.POST("/", h.CreateDelivery)
}