package location

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const driverGeoKey = "driver_locations"

type Store interface {
	SetDriverLocation(ctx context.Context, driverID string, lat, lng float64) error
	GetNearbyDrivers(ctx context.Context, lat, lng float64, radius float64) ([]string, error)
	DeleteDriver(ctx context.Context, driverID string) error
}

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) Store {
	return &RedisStore{
		client: client,
	}
}

func (r *RedisStore) SetDriverLocation(ctx context.Context, driverID string, lat, lng float64) error {
	return r.client.GeoAdd(ctx,driverGeoKey, &redis.GeoLocation{
		Name: driverID,
		Longitude: lng,
		Latitude: lat,	
	}).Err()
}

func (r *RedisStore) GetNearbyDrivers(ctx context.Context, lat, lng float64, radius float64) ([]string, error) {
	locations, err := r.client.GeoSearch(ctx,driverGeoKey, &redis.GeoSearchQuery{
	Latitude:   lat,
        Longitude:  lng,
        Radius:     radius,
        RadiusUnit: "km",
        Sort:       "ASC",
        Count:      20,	
	}).Result()
	if err != nil {
		return nil, err
	}

	return locations, nil
}

func (r *RedisStore) DeleteDriver(ctx context.Context, driverID string) error {
    return r.client.ZRem(ctx, driverGeoKey, driverID).Err()
}