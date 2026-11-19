package repository

import (
	"context"

	"github.com/rbx0o/go-todo-api/internal/core/database"
	"go.uber.org/zap"
)

type taskRepo struct {
	DB     *database.DataBase
	logger *zap.Logger
}

func NewTaskRepo(db *database.DataBase, logger *zap.Logger) *taskRepo {
	return &taskRepo{
		DB:     db,
		logger: logger,
	}
}

func (t *taskRepo) GetAll(ctx context.Context) {
	sqlQuery := `
	SELECT id, name, description, createdTime, isDone, finishedTime
	FROM public.tasks
	`

	rows, err := t.DB.Conn.Query(ctx, sqlQuery)
	if err != nil {
		t.logger.Error(
			"SQL query error",
			zap.Error(err),
		)
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan()
	}

}
