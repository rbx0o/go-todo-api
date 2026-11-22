package handlers

import (
	"net/http"

	"github.com/rbx0o/go-todo-api/internal/tasks/service"
	"go.uber.org/zap"
)

type HandlerInterface interface {
	GetMux() *http.ServeMux
}

type handler struct {
	service service.TaskServiceInterface
	mux     *http.ServeMux
	logger  *zap.Logger
}

func New(service service.TaskServiceInterface, logger *zap.Logger) *handler {
	mux := http.NewServeMux()

	h := &handler{
		service: service,
		mux:     mux,
		logger:  logger,
	}

	h.mux.HandleFunc("GET /tasks", h.GetAll)

	return h
}

func (h *handler) GetMux() *http.ServeMux {
	return h.mux
}
