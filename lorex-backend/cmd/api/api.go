package main

import (
	"time"
	"net/http"
	"log/slog"
	"os"
	"os/signal"
	"context"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
  router := gin.Default()
  
  router.GET("/health", func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"status": "healthy",
	})
  })

  return router
}


func StartServer(router *gin.Engine, app *App) {
	
	srv := &http.Server{
		Addr: ":" + app.Config.Port,
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 10 * time.Second,
	}

	//graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("listen", "error", err)
		}
	}()

	slog.Info("Server started", "port", app.Config.Port)

	// Wait for interrupt signal to gracefully shutdown the server
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit
    slog.Info("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        slog.Error("Server forced to shutdown", "error", err)
    }

	app.DBConn.Close()

    slog.Info("Server exiting")
}