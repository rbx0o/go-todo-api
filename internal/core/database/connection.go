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

func CreateConnection(ctx context.Context, logger *zap.Logger) (*DataBase, <-chan error, error) {
	errCh := make(chan error, 1)
	conn, err := pgx.Connect(ctx, os.Getenv("CONN_STRING"))

	logger.Info("Connection to database created successfully")

	if err != nil {
		logger.Error(
			"Error connect to database",
			zap.Error(err),
		)
		close(errCh)
		return &DataBase{}, errCh, err
	}

	if err = conn.Ping(ctx); err != nil {
		logger.Error(
			"Error ping database",
			zap.Error(err),
		)
		close(errCh)
		return &DataBase{}, errCh, err
	}

	go func() {
		<-ctx.Done()

		closeDbCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
			close(errCh)
		}()

		if err := conn.Close(closeDbCtx); err != nil {
			logger.Error(
				"The attempt to close the connection failed",
				zap.Error(err),
			)
			errCh <- err
			return
		}

		logger.Info("Connection closed successfully")
	}()

	return &DataBase{
		ctx:    ctx,
		conn:   conn,
		logger: logger,
	}, errCh, nil
}
