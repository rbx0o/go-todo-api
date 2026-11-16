package database

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type DataBase struct {
	ctx    context.Context
	conn   *pgx.Conn
	logger *zap.Logger
}

func CreateConnection(ctx context.Context, logger *zap.Logger) (*DataBase, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("CONN_STRING"))
	if err != nil {
		logger.Error(
			"Error connect to database",
			zap.Error(err),
		)
		return &DataBase{}, err
	}

	if err = conn.Ping(ctx); err != nil {
		logger.Error(
			"Error ping database",
			zap.Error(err),
		)
		return &DataBase{}, err
	}

	go func() {
		<-ctx.Done()

		closeDbCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := conn.Close(closeDbCtx); err != nil {
			logger.Error(
				"The attempt to close the connection failed",
				zap.Error(err),
			)
			return
		}

		logger.Info("Connection closed successfully")
	}()

	return &DataBase{
		ctx:    ctx,
		conn:   conn,
		logger: logger,
	}, nil
}
