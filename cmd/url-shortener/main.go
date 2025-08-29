package main

import (
	"context"
	"log/slog"
	"os"
	"time"
	"url-shortener/internal/config"
	"url-shortener/internal/repository"
	"url-shortener/internal/routes"
	"url-shortener/internal/service"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("start of work", slog.String("env", cfg.Env))

	repo, err := repository.NewRepository(ctx, cfg)
	if err != nil {
		log.Error("failed connection with")
		os.Exit(1)
	}
	defer repo.Close()
	log.Info("connect with database")

	svc := service.NewURLService(repo)

	r := routes.SetupRouter(svc, log)

	if err := r.Run(":8080"); err != nil {
		log.Error("failed to start server", slog.Any("err", err))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return log
}
