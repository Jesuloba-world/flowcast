package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jesuloba-world/flowcast/internal/config"
	"github.com/Jesuloba-world/flowcast/internal/infrastructure/database"
	"github.com/Jesuloba-world/flowcast/internal/infrastructure/dragonfly"
	"github.com/Jesuloba-world/flowcast/internal/logger"
	"github.com/Jesuloba-world/flowcast/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	logger := logger.New(cfg.Logging)

	db, err := database.New(*cfg)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	dragonfly, err := dragonfly.New(cfg.Dragonfly)
	if err != nil {
		logger.Error("Failed to connect to DragonflyDB", "error", err)
		os.Exit(1)
	}
	defer dragonfly.Close()

	// Start the server
	srv := server.New(cfg, db, dragonfly.Client, logger)

	httpServer := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: srv.Router(),
	}

	// graceful shutdown
	go func() {
		logger.Info("starting server", "port", cfg.Server.Port, "environment", cfg.Server.Environment)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	// wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server...")

	// gracefully shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error("server forced to shutdown", "error", err)
		os.Exit(1)
	}

	logger.Info("Server exited gracefully")
}
