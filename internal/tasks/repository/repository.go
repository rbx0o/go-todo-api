package repository

import (
	"context"
	"time"

	"github.com/rbx0o/go-todo-api/internal/core/database"
	"github.com/rbx0o/go-todo-api/internal/core/domain"
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

func (repo *taskRepo) GetAll(ctx context.Context) (map[int64]domain.Task, error) {
	sqlQuery := `
	SELECT id, name, description, created_time, is_done, finished_time
	FROM public.tasks
	`

	rows, err := repo.DB.Conn.Query(ctx, sqlQuery)
	if err != nil {
		repo.logger.Error(
			"SQL query error",
			zap.Error(err),
		)
		return nil, err
	}
	defer rows.Close()

	tasks := make(map[int64]domain.Task, 0)

	for rows.Next() {
		var (
			id           int64
			name         string
			description  string
			createdTime  time.Time
			isDone       bool
			finishedTime *time.Time
		)
		rows.Scan(&id, &name, &description, &createdTime, &isDone, &finishedTime)
		task := domain.NewTask(
			id,
			name,
			description,
			createdTime,
			isDone,
			finishedTime,
		)
		tasks[id] = *task
	}

	if err := rows.Err(); err != nil {
		repo.logger.Error(
			"SQL query error",
			zap.Error(err),
		)
		return nil, err
	}

	return tasks, nil
}
