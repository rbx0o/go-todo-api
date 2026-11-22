package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func (h *handler) GetAll(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	tasks, err := h.service.GetAll(ctx)
	if err != nil {
		h.logger.Error(
			"GET /tasks handler error from service.GetAll",
			zap.Error(err),
		)
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		h.logger.Error(
			"GET /tasks handler error from json.MarshalIndent",
			zap.Error(err),
		)
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write(b)
	if err != nil {
		h.logger.Error(
			"GET /tasks handler error from response.Write",
			zap.Error(err),
		)
	}
}
