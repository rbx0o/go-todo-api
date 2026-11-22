package database

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type DataBase struct {
	Conn   *pgx.Conn
	logger *zap.Logger
}

func CreateConnection(logger *zap.Logger) (*DataBase, <-chan error, error) {
	errCh := make(chan error, 1)
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()
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

	return &DataBase{
		Conn:   conn,
		logger: logger,
	}, errCh, nil
}

func (db *DataBase) Stop() error {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	if !db.Conn.IsClosed() {
		err := db.Conn.Close(ctx)
		if err != nil {
			db.logger.Error(
				"DataBase connection close error",
				zap.Error(err),
			)
			return err
		}

		db.logger.Info("DataBase stopped successfully")
		return nil
	}

	db.logger.Warn("DataBase connection already closed")
	return errors.New("DataBase connection already closed")
}
