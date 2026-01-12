package main

import (
	"log/slog"

	"github.com/kubatov-hub/url-shortener-golang/internal/config"
	"github.com/kubatov-hub/url-shortener-golang/internal/logger"
	"github.com/kubatov-hub/url-shortener-golang/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// TODO: init storage: sqlite
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", slog.String("error", err.Error()))
		return
	}

	_ = storage

	// TODO: init router: chi, "chi render"

	// TODO: run server:
}
