package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

type DataBase struct {
	ctx  context.Context
	conn pgx.Conn
}

func CreateConnection(ctx context.Context) (*DataBase, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("CONN_STRING"))
	if err != nil {
		return &DataBase{}, err
	}

	if err = conn.Ping(ctx); err != nil {
		return &DataBase{}, err
	}

	return &DataBase{
		ctx:  ctx,
		conn: *conn,
	}, nil
}
