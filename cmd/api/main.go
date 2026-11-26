package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"

	"mypulse/internal/attendance"
	"mypulse/internal/platform/config"
	"mypulse/internal/platform/database"
	"mypulse/internal/platform/httpx"
)

func main() {
	if err := run(); err != nil {
		slog.Error("server berhenti", "error", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pool, err := database.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer pool.Close()

	// Dependency injection manual: repository -> service -> handler
	attendanceHandler := attendance.NewHandler(
		attendance.NewService(attendance.NewRepository(pool)),
	)

	app := fiber.New(fiber.Config{
		ErrorHandler: httpx.ErrorHandler,
	})
	app.Use(recoverer.New(), requestid.New(), logger.New())

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	api := app.Group("/api/v1")
	attendanceHandler.Routes(api.Group("/attendance"))

	errCh := make(chan error, 1)
	go func() {
		slog.Info("server jalan", "addr", cfg.HTTPAddr)
		errCh <- app.Listen(cfg.HTTPAddr)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return app.ShutdownWithContext(shutdownCtx)
}
