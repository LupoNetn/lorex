package location

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luponetn/lorex/utils"
)

type Handler struct {
	store Store
}

func NewHandler(store Store) *Handler {
	return &Handler{
		store: store,
	}
}

type UpdateLocationRequest struct {
	Lat float64 `json:"lat" binding:"required"`
	Lng float64 `json:"lng" binding:"required"`
}

func (h *Handler) UpdateLocation(c *gin.Context) {
	var req UpdateLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid or missing lat/lng",
		})
		slog.Error("failed to bind location update request", "error", err)
		return
	}

	val, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	claims, ok := val.(*utils.CustomClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse claims"})
		return
	}

	driverID := claims.ID

	err := h.store.SetDriverLocation(c.Request.Context(), driverID, req.Lat, req.Lng)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update location",
		})
		slog.Error("failed to update driver location", "error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "location updated successfully",
	})
}
