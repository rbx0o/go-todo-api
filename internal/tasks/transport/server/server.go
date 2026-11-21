package server

import (
	"errors"
	"net/http"

	"go.uber.org/zap"
)

type httpServer struct {
	addr    string
	handler *http.Handler
	logger  *zap.Logger
}

func New(addr string, handler *http.Handler, logger *zap.Logger) *httpServer {
	return &httpServer{
		addr:    addr,
		handler: handler,
		logger:  logger,
	}
}

func (server *httpServer) Start() <-chan error {
	chErr := make(chan error, 1)

	go func() {
		err := http.ListenAndServe(server.addr, *server.handler)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			chErr <- err
		}
		close(chErr)
	}()

	return chErr
}
