package main

import (
	"log/slog"

	"github.com/kubatov-hub/url-shortener-golang/internal/config"
	"github.com/kubatov-hub/url-shortener-golang/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server:
}
