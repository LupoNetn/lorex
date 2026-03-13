package delivery

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luponetn/lorex/internal/db/sqlc"
	"github.com/luponetn/lorex/utils"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

//implement delivery functionality
func (h *Handler) CreateDelivery(c *gin.Context) {
	var req CreateDeliveryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utils.FormatValidationError(err),
		})
		slog.Error(err.Error())
		return
	}

	params := sqlc.CreateDeliveryParams{
		CustomerID:      req.CustomerID,
		PickupAddress:   req.PickupAddress,
		DeliveryAddress: req.DeliveryAddress,
		PickupLat:       req.PickupLat,
		PickupLng:       req.PickupLng,
		DeliveryLat:     req.DeliveryLat,
		DeliveryLng:     req.DeliveryLng,
		PackageType:     req.PackageType,
		Weight:          req.Weight,
		Price:           req.Price,
	}

	delivery, err := h.service.CreateDelivery(c.Request.Context(),params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong failed to create delivery",
			"error": err.Error(),
		})
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delivery created successfully",
		"data": delivery,
	})
}