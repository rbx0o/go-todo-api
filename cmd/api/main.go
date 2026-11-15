package main

import (
	"context"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/k0kubun/pp/v3"
	"github.com/rbx0o/go-todo-api/internal/core/database"
)

func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())

	db, err := database.CreateConnection(ctx)
	if err != nil {
		fmt.Println(err)
	}

	pp.Println(db)

	ctxCancel()
}
