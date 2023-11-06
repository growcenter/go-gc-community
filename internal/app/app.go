package app

import (
	"context"
	"errors"
	"go-gc-community/internal/config"
	handler "go-gc-community/internal/delivery/http"
	"go-gc-community/internal/repositories"
	"go-gc-community/internal/server"
	"go-gc-community/internal/usecases"
	"go-gc-community/pkg/database/msql"
	"go-gc-community/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		logger.Error(err)
	}

	// Database
	msql, err := msql.Connect(cfg.MSql.User, cfg.MSql.Password, cfg.MSql.Host, cfg.MSql.Name)
	if err != nil {
		logger.Error(err)
	}

	repository := repositories.NewRepositories(msql)
	usecase := usecases.NewUsecases(usecases.Dependencies{
		Repository: repository,
	})
	handler := handler.NewHandler(usecase)
	server := server.NewServer(cfg, handler.Init(cfg))

	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := server.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}