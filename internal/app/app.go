package app

import (
	"context"
	"errors"
	"fmt"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rbx0o/go-todo-api/internal/core/apperr"
	"github.com/rbx0o/go-todo-api/internal/core/database"
	"github.com/rbx0o/go-todo-api/internal/core/logger"
	"github.com/rbx0o/go-todo-api/internal/tasks/repository"
	"github.com/rbx0o/go-todo-api/internal/tasks/service"
	"github.com/rbx0o/go-todo-api/internal/tasks/transport/handlers"
	"github.com/rbx0o/go-todo-api/internal/tasks/transport/server"
)

func Run() {
	logger := logger.New()
	logger.Info("logger started successfully")
	sigCtx, sigCtxStop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		err := logger.Sync()
		if err != nil {
			err = fmt.Errorf("app stop %w", err)
			logger.Error(err.Error())
		}
		sigCtxStop()
	}()

	db, err := database.CreateConnection(logger)

	if err != nil {
		err = fmt.Errorf("app run: %w", err)
		logger.Error(err.Error())
		return
	}

	repo := repository.NewTaskRepo(db, logger)
	service := service.NewTaskService(repo, logger)
	handler := handlers.New(service, logger)
	server := server.New("localhost:8080", handler, logger)
	errCh := server.Start()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		<-sigCtx.Done()

		err := db.Stop()
		err = fmt.Errorf("app stop: %w", err)
		if err != nil {
			if errors.Is(err, apperr.DBConnectionAlreadyClosed) {
				logger.Warn(err.Error())
			} else {
				logger.Error(err.Error())
			}
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		<-errCh
		wg.Done()
	}()

	wg.Wait()
}
