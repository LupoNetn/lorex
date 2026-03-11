package logger 

import (
	"log/slog"
	"os"
)

var log *slog.Logger


func Init() {
    handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	log = slog.New(handler)

	slog.SetDefault(log)
}