package main

import (
	"log/slog"
	"os"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luponetn/lorex/internal/auth"
	"github.com/luponetn/lorex/internal/config"
	"github.com/luponetn/lorex/internal/db"
	"github.com/luponetn/lorex/internal/db/sqlc"
	"github.com/luponetn/lorex/internal/delivery"
	"github.com/luponetn/lorex/internal/logger"
	"github.com/luponetn/lorex/internal/store"
	"github.com/luponetn/lorex/internal/tasks"
)

type App struct {
	Config *config.Config
	DBConn *pgxpool.Pool
}


func main() {
	logger.Init()

	//load config
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("unable to load config for this app", "error", err)
		return
	}


	//load database
	dbConn, err := db.ConnectDB(cfg)
	if err != nil {
		slog.Error("could not connect to db", "error", err)
		return
	}

	app := &App{
		Config: cfg,
		DBConn: dbConn,
	}

	//setup server 
	router := SetUpRouter()

	//setup asynq client
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		slog.Error("redis url is not set")
		return
	}

	redisOpt, err := asynq.ParseRedisURI(redisURL)
	if err != nil {
		slog.Error("failed to parse redis uri")
		return
	}

	client := asynq.NewClient(redisOpt)
	defer client.Close()
    
	// setup asynq enquer with the client
	enquer := tasks.NewAsynqClient(client)

	// load store
	q := sqlc.New(dbConn) // This might error if driver mismatch, but follows your requested flow

	// setup asynq server (The Worker)
	taskProcessor := tasks.NewRedisTaskProcessor(redisOpt, q)
	go func() {
		slog.Info("starting asynq task processor")
		if err := taskProcessor.Start(); err != nil {
			slog.Error("failed to start task processor", "error", err)
		}
	}()

	AuthPGStore := store.NewAuthPostgresStore(q)
	DeliveryPGStore := store.NewDeliveryPostgresStore(q)

	// hook up handlers,routers and services
	authService := auth.NewSvc(AuthPGStore)
	authHandler := auth.NewHandler(authService)
	auth.RegisterRoutes(router, authHandler)

	// delivery
	deliveryService := delivery.NewSvc(DeliveryPGStore, enquer)
	deliveryHandler := delivery.NewHandler(deliveryService)
	delivery.RegisterRoutes(router, deliveryHandler)

	StartServer(router, app)
}