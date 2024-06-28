package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/celsopires1999/estimation/internal/usecase"
)

type ProjectHandler struct {
	createProjectUseCase *usecase.CreateProjectUseCase
}

func NewProjectsHandler(createProjectUseCase *usecase.CreateProjectUseCase) *ProjectHandler {
	return &ProjectHandler{
		createProjectUseCase: createProjectUseCase,
	}
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProjectInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if errors := ValidatePayload(input); errors != nil {
		WriteValidationError(w, http.StatusBadRequest, errors)
		return
	}

	output, err := h.createProjectUseCase.Execute(context.Background(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusCreated, output)
}
