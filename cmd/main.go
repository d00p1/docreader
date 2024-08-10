package main

import (
	"log/slog"

	"github.com/d00p1/docreader/config"
)

func main() {
	cfg := config.Load()

	log := slog.Default()

	log.Info("logger system complite", slog.String("env", cfg.Env))
}
