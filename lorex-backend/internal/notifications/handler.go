package notifications

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetNotifications(c *gin.Context) {
	userVal, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user := userVal.(struct {
		ID    string
		Email string
		Role  string
	})

	var receiverID pgtype.UUID
	if err := receiverID.Scan(user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse user ID"})
		return
	}

	notifications, err := h.service.GetNotifications(c.Request.Context(), receiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch notifications"})
		slog.Error("failed to fetch notifications", "error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": notifications,
	})
}

func (h *Handler) MarkAsRead(c *gin.Context) {
	idStr := c.Param("id")
	var id pgtype.UUID
	if err := id.Scan(idStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid notification ID"})
		return
	}

	if err := h.service.MarkAsRead(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to mark notification as read"})
		slog.Error("failed to mark notification as read", "error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "notification marked as read",
	})
}
