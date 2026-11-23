package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rbx0o/go-todo-api/internal/core/apperr"
	"go.uber.org/zap"
)

type DataBase struct {
	Conn   *pgx.Conn
	logger *zap.Logger
}

func CreateConnection(logger *zap.Logger) (*DataBase, error) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()
	conn, err := pgx.Connect(ctx, os.Getenv("CONN_STRING"))

	if err != nil {
		err = fmt.Errorf("connect to PostgreSQL:  %w", err)
		return &DataBase{}, err
	}

	logger.Info("connection to database created successfully")

	if err = conn.Ping(ctx); err != nil {
		err = fmt.Errorf("ping database: %w", err)
		return &DataBase{}, err
	}

	return &DataBase{
		Conn:   conn,
		logger: logger,
	}, nil
}

func (db *DataBase) Stop() error {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	if !db.Conn.IsClosed() {
		err := db.Conn.Close(ctx)
		if err != nil {
			err = fmt.Errorf("database connection close: %w", err)
			return err
		}

		db.logger.Info("database stopped successfully")
		return nil
	} else {
		return apperr.DBConnectionAlreadyClosed
	}
}
