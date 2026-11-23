package server

import (
	"errors"
	"net/http"

	"github.com/rbx0o/go-todo-api/internal/tasks/transport/handlers"
	"go.uber.org/zap"
)

type httpServer struct {
	handler handlers.HandlerInterface
	server  *http.Server
	logger  *zap.Logger
}

func New(addr string, handler handlers.HandlerInterface, logger *zap.Logger) *httpServer {
	server := http.Server{
		Addr:    addr,
		Handler: handler.GetMux(),
	}

	return &httpServer{
		handler: handler,
		server:  &server,
		logger:  logger,
	}
}

func (server *httpServer) Start() <-chan error {
	chErr := make(chan error, 1)

	go func() {
		err := server.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			chErr <- err
		}
		close(chErr)
	}()

	return chErr
}
