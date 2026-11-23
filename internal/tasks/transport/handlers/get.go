package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rbx0o/go-todo-api/internal/tasks/transport/dto"
	"go.uber.org/zap"
)

func (h *handler) GetAll(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	tmpTasks, err := h.service.GetAll(ctx)
	if err != nil {
		h.logger.Error(
			"GET /tasks handler error from service.GetAll",
			zap.Error(err),
		)
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	tasks := make(map[int64]dto.Task, len(tmpTasks))
	for key, value := range tmpTasks {
		tmpTaskCopy := value.Get()
		tmpDtoTask := dto.Task{
			ID:           tmpTaskCopy.ID,
			Name:         tmpTaskCopy.Name,
			Description:  tmpTaskCopy.Description,
			CreatedTime:  tmpTaskCopy.CreatedTime,
			IsDone:       tmpTaskCopy.IsDone,
			FinishedTime: tmpTaskCopy.FinishedTime,
		}
		tasks[key] = tmpDtoTask
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
