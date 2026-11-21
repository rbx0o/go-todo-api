package handlers

import (
	"net/http"

	"github.com/rbx0o/go-todo-api/internal/tasks/service"
)

type handler struct {
	service service.TaskServiceInterface
	mux     *http.ServeMux
}

func New(service service.TaskServiceInterface) *handler {
	mux := http.NewServeMux()

	return &handler{
		service: service,
		mux:     mux,
	}
}
