package notifications

import (
	"github.com/gin-gonic/gin"
	"github.com/luponetn/lorex/internal/middleware"
)

func RegisterRoutes(router *gin.Engine, h *Handler) {
	notificationGroup := router.Group("/notifications")
	notificationGroup.Use(middleware.AuthMiddleware())

	notificationGroup.GET("/", h.GetNotifications)
	notificationGroup.PUT("/:id/read", h.MarkAsRead)
}
