package main

import (
	"context"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rbx0o/go-todo-api/internal/core/database"
)

func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())

	_, err := database.New(ctx)
	if err != nil {
		fmt.Println(err)
	}

	ctxCancel()
}
