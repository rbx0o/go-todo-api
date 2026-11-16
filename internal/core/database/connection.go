package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

type DataBase struct {
	ctx  context.Context
	conn *pgx.Conn
}

func New(ctx context.Context) (*DataBase, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("CONN_STRING"))
	if err != nil {
		return &DataBase{}, err
	}

	if err = conn.Ping(ctx); err != nil {
		return &DataBase{}, err
	}

	go func() {
		<-ctx.Done()

		closeDbCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := conn.Close(closeDbCtx); err != nil {
			fmt.Println("Попытка закрытия подключения завершилась неудачно: ", err)
			return
		}

		fmt.Println("Подключение успешно закрыто")
	}()

	return &DataBase{
		ctx:  ctx,
		conn: conn,
	}, nil
}
