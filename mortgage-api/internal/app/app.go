package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/artomsopun/mortgage/mortgage-api/internal/config"
	"github.com/artomsopun/mortgage/mortgage-api/internal/delivery"
	"github.com/artomsopun/mortgage/mortgage-api/internal/repository"
	"github.com/artomsopun/mortgage/mortgage-api/internal/server"
	"github.com/artomsopun/mortgage/mortgage-api/internal/service"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/auth"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/database"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/hash"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/logger"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)

		return
	}

	// Dependencies
	db := database.NewDB(cfg.MySql.User, cfg.MySql.Password, cfg.MySql.Host, cfg.MySql.Port, cfg.MySql.Name)
	if err != nil {
		logger.Error(err)

		return
	}

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)

	authManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		logger.Error(err)

		return
	}

	// Services, Repos & API Handlers
	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos:           repos,
		Hasher:          hasher,
		AuthManager:     authManager,
		AccessTokenTTL:  cfg.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL: cfg.Auth.JWT.RefreshTokenTTL,
	})
	handlers := delivery.NewHandler(services, authManager)

	// HTTP Server
	srv := server.NewServer(cfg, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	if sqlDB, err := db.DB(); err != nil {
		logger.Error(err.Error())
	} else {
		if err := sqlDB.Close(); err != nil {
			logger.Error(err.Error())
		}
	}
}
