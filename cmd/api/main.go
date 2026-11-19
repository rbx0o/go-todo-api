package main

import (
	"context"
	"sync"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rbx0o/go-todo-api/internal/core/database"
	"github.com/rbx0o/go-todo-api/internal/core/logger"
)

func main() {
	logger := logger.New()
	logger.Info("Logger started successfully")
	ctx, ctxCancel := context.WithCancel(context.Background())

	_, errCh, _ := database.CreateConnection(ctx, logger)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		<-errCh
		wg.Done()
	}()

	ctxCancel()
	wg.Wait()
}
