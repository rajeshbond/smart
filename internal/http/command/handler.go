package command

import (
	"encoding/json"
	"fmt"
	"net/http"

	dto "github.com/rajeshbond/smart/internal/mqtt/assembly/production_dto"
)

type Handler struct {
	commandService *CommandService
}

func NewHandler(commandService *CommandService) *Handler {
	return &Handler{commandService: commandService}
}

func (h *Handler) ResetCounterHandler(
	w http.ResponseWriter, r *http.Request,
) {
	var req dto.ResetCounterRequest

	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}
	fmt.Println(req)
	if err := h.commandService.ResetCounter(ctx, req); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return

	}
}
