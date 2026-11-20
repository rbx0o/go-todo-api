package service

import (
	"context"

	"github.com/rbx0o/go-todo-api/internal/core/domain"
	"github.com/rbx0o/go-todo-api/internal/tasks/repository"
	"go.uber.org/zap"
)

type taskService struct {
	repo   repository.TaskRepoInterface
	logger *zap.Logger
}

func NewTaskService(repo repository.TaskRepoInterface, logger *zap.Logger) *taskService {
	return &taskService{
		repo:   repo,
		logger: logger,
	}
}

func (service *taskService) GetAll(ctx context.Context) (map[int64]domain.Task, error) {
	return service.repo.GetAll(ctx)
}
