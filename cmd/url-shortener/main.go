package main

import (
	"golang.org/x/exp/slog"
	"os"
	"rest_api_project/internal/config"
	"rest_api_project/internal/lib/logger/sl"
	"rest_api_project/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage
	id, err := storage.SaveUrl("https://www.youtube.com/watch?v=rCJvW2xgnk0&t=70s", "youtube")
	if err != nil {
		log.Error("failed to save url", sl.Err(err))
		os.Exit(1)
	}

	log.Info("save url :", slog.Int64("id", id))

	id, err = storage.SaveUrl("https://www.youtube.com/watch?v=rCJvW2xgnk0&t=70s", "youtube")
	if err != nil {
		log.Error("failed to save url", sl.Err(err))
		os.Exit(1)
	}
}
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	}

	return log
}
