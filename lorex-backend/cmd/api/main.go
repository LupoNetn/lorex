package main

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luponetn/lorex/internal/config"
	"github.com/luponetn/lorex/internal/db"
	"github.com/luponetn/lorex/internal/logger"
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

	StartServer(router, app)
}