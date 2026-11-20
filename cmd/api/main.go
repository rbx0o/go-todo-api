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
	defer logger.Sync()
	dbConnectionCtx, dbConnectionCtxCancel := context.WithCancel(context.Background())

	_, errCh, _ := database.CreateConnection(dbConnectionCtx, logger)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		<-errCh
		wg.Done()
	}()

	dbConnectionCtxCancel()
	wg.Wait()
}
