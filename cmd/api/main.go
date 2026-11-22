package main

import (
	"sync"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rbx0o/go-todo-api/internal/core/database"
	"github.com/rbx0o/go-todo-api/internal/core/logger"
	"github.com/rbx0o/go-todo-api/internal/tasks/repository"
	"github.com/rbx0o/go-todo-api/internal/tasks/service"
	"github.com/rbx0o/go-todo-api/internal/tasks/transport/handlers"
	"github.com/rbx0o/go-todo-api/internal/tasks/transport/server"
)

func main() {
	logger := logger.New()
	logger.Info("Logger started successfully")
	defer logger.Sync()

	db, errCh, _ := database.CreateConnection(logger)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		<-errCh
		wg.Done()
	}()

	repo := repository.NewTaskRepo(db, logger)
	service := service.NewTaskService(repo, logger)
	handler := handlers.New(service, logger)
	server := server.New("localhost:8080", handler, logger)
	chErr := server.Start()

	wg.Add(1)
	go func() {
		<-chErr
		wg.Done()
	}()

	wg.Wait()
}
