package location

import (
	"context"
	"time"
	"log/slog"

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
	err := r.client.GeoAdd(ctx,driverGeoKey, &redis.GeoLocation{
		Name: driverID,
		Longitude: lng,
		Latitude: lat,	
	}).Err()
	if err != nil {
		return err
	}

	return r.client.Set(ctx,"active:"+driverID,"true", 30*time.Second).Err()
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

	//verify sticky notes only keep drivers with active heartbeats
	var activeDrivers []string
	for _, driverID := range locations {
		exists, err := r.client.Exists(ctx, "active:"+driverID).Result()
		if err != nil {
			slog.Error("error checking driver activity", "error", err)
			continue
		}
		if exists == 0 {
			// sticky note gone — clean up the dot too
            r.client.ZRem(ctx, "drivers:active", driverID)
            continue
		}
		activeDrivers = append(activeDrivers, driverID)
	}

	return activeDrivers, nil
}

func (r *RedisStore) DeleteDriver(ctx context.Context, driverID string) error {
    return r.client.ZRem(ctx, driverGeoKey, driverID).Err()
}