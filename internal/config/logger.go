package config

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	Logger = slog.New(handler)
}
