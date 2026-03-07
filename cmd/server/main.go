package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MediStatTech/biometric-service/internal"
	"github.com/MediStatTech/biometric-service/internal/app"
	"github.com/MediStatTech/biometric-service/internal/cron/registry"
	"github.com/MediStatTech/biometric-service/internal/health"
	"github.com/MediStatTech/biometric-service/pkg"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("error loading .env file: %w", err))
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	pkgInstance, err := pkg.New(ctx)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize PKG: %v\n", err))
	}

	appInstance, err := app.New(pkgInstance)
	if err != nil {
		pkgInstance.Logger.Fatal("Failed to create app instance.", map[string]any{
			"error": err,
		})
		return
	}

	grpcServer, err := internal.New(ctx, pkgInstance, appInstance)
	if err != nil {
		pkgInstance.Logger.Fatal("Failed to create gRPC server.", map[string]any{
			"error": err,
		})
		return
	}

	// Start HTTP health server for Kubernetes probes
	pkgInstance.Logger.Info("Starting health server for Kubernetes probes...", map[string]any{})
	healthServer := health.NewHealthServer(pkgInstance.Logger, ":8088")
	if err := healthServer.Start(); err != nil {
		pkgInstance.Logger.Fatal("Failed to start health server", map[string]any{
			"error": err,
		})
		return
	}
	pkgInstance.Logger.Info("Health server started successfully", map[string]any{
		"port": "8080",
	})

	pkgInstance.Logger.Info("Starting gRPC server", map[string]any{
		"service": "biometric",
		"address": grpcServer.Address(),
	})

	cronRegistry := registry.New(&registry.Options{
		Log:     pkgInstance.Logger,
		Process: appInstance.Biometric.Process,
	})

	cronRegistry.Start(ctx)

	go func() {
		defer cancel()
		if err := grpcServer.Serve(); err != nil {
			pkgInstance.Logger.Fatal("gRPC server error", map[string]any{
				"error": err,
			})
		}
	}()

	<-ctx.Done()

	pkgInstance.Logger.Info("Shutting down server", map[string]any{})

	// Shutdown health server
	pkgInstance.Logger.Info("Shutting down health server...", map[string]any{})
	if err := healthServer.Shutdown(context.Background()); err != nil {
		pkgInstance.Logger.Error("Error during health server shutdown", map[string]any{
			"error": err,
		})
	} else {
		pkgInstance.Logger.Info("Health server shutdown complete", map[string]any{})
	}

	// ---- Cleanup ----
	cronRegistry.Shutdown()

	if err := grpcServer.Shutdown(context.Background()); err != nil {
		pkgInstance.Logger.Error("gRPC server shutdown error", map[string]any{
			"error": err,
		})
	}
}
