package app

import (
	"context"
	"errors"
	"fmt"
	"go-gc-community/internal/config"
	handler "go-gc-community/internal/delivery/http"
	"go-gc-community/internal/repositories"
	"go-gc-community/internal/server"
	"go-gc-community/internal/usecases"
	"go-gc-community/pkg/authorization"
	"go-gc-community/pkg/database/msql"
	"go-gc-community/pkg/google"
	"go-gc-community/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		//logger.Error(err)
		logger.Logger.Error("Config error", zap.Error(err))
	}

	// Database
	msql, err := msql.Connect(cfg.MSql.User, cfg.MSql.Password, cfg.MSql.Host, cfg.MSql.Name)
	if err != nil {
		//logger.Error(err)
		logger.Logger.Error("Database error", zap.Error(err))
	}

	// Token
	authService, err := authorization.NewAuthorization(cfg.Auth.Secret, cfg.Auth.TokenExpiry)
	if err != nil {
		//logger.Error(err)
		logger.Logger.Error("Auth error", zap.Error(err))
	}

	// Google
	authGoogle, err := google.NewGoogle(cfg.Google.State, cfg.Google.ClientId, cfg.Google.ClientSecret, cfg.Google.RedirectUrl)
	if err != nil {
		//logger.Error(err)
		logger.Logger.Error("Google Oauth error", zap.Error(err))
	}

	repository := repositories.NewRepositories(msql)
	usecase := usecases.NewUsecases(usecases.Dependencies{
		Repository: repository,
		Authorization: authService,
		Google: authGoogle,
	})
	handler := handler.NewHandler(usecase, authService)
	server := server.NewServer(cfg, handler.Init(cfg))

	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			//logger.Errorf("error occurred while running http server: %s\n", err.Error())
			logger.Logger.Error(fmt.Sprintf("error occurred while running http server: %s\n", err), zap.Error(err))
		}
	}()

	//logger.Info("Server Started")
	logger.Logger.Info("Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := server.Stop(ctx); err != nil {
		//logger.Errorf("failed to stop server: %v", err)
		logger.Logger.Error(fmt.Sprintf("Failed to stop server: %s\n", err), zap.Error(err))
	}
}