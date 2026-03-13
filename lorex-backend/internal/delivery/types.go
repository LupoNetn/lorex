package delivery

import "github.com/jackc/pgx/v5/pgtype"

type CreateDeliveryRequest struct {
	CustomerID      pgtype.UUID    `json:"customer_id" binding:"required"`
	PickupAddress   string         `json:"pickup_address" binding:"required"`
	DeliveryAddress string         `json:"delivery_address" binding:"required"`
	PickupLat       pgtype.Float8  `json:"pickup_lat" binding:"required"`
	PickupLng       pgtype.Float8  `json:"pickup_lng" binding:"required"`
	DeliveryLat     pgtype.Float8  `json:"delivery_lat" binding:"required"`
	DeliveryLng     pgtype.Float8  `json:"delivery_lng" binding:"required"`
	PackageType     pgtype.Text    `json:"package_type"`
	Weight          pgtype.Float8  `json:"weight"`
	Price           pgtype.Numeric `json:"price"`
}