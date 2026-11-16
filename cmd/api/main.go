package main

import (
	"context"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rbx0o/go-todo-api/internal/core/database"
	"github.com/rbx0o/go-todo-api/internal/core/logger"
)

func main() {
	logger := logger.New()
	logger.Info("Logger started successfully")
	ctx, ctxCancel := context.WithCancel(context.Background())

	database.CreateConnection(ctx, logger)

	//pp.Println(db)

	ctxCancel()

	time.Sleep(3 * time.Second)
}
