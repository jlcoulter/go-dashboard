package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"log/slog"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/jlcoulter/go-dashboard-template/internal/config"
	"github.com/jlcoulter/go-dashboard-template/internal/handler"
	"github.com/jlcoulter/go-dashboard-template/internal/logging"
	"github.com/jlcoulter/go-dashboard-template/internal/template"
)

func main() {
	cfg := config.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: parseLogLevel(cfg.LogLevel),
	}))
	slog.SetDefault(logger)

	// Load templates
	tmpl, err := template.New()
	if err != nil {
		slog.Error("failed to load templates", "error", err)
		os.Exit(1)
	}

	r := chi.NewRouter()
	r.Use(logging.RequestLogger(logger))

	// Static files
	r.Handle("/static/*", handler.StaticRoutes())

	// Health endpoints
	r.Get("/healthz", handler.Healthz)
	r.Get("/readyz", handler.Readyz)

	// HTML pages
	r.Get("/", handler.Dashboard(tmpl))
	r.Get("/dashboard", handler.Dashboard(tmpl))

	// HTMX partials
	r.Get("/partials/stats", handler.StatsPartial(tmpl))

	// JSON API
	r.Route("/api", func(r chi.Router) {
		r.Get("/stats", handler.StatsAPI)
	})

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		slog.Info("starting server", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("server forced to shutdown", "error", err)
	}
	slog.Info("server stopped")
}

func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}