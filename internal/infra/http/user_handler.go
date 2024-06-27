package http

import (
	"encoding/json"
	"net/http"

	"github.com/celsopires1999/estimation/internal/service"
)

type PensonsHandler struct {
	service *service.UserService
}

func NewPersonsHandler(personService *service.UserService) *PensonsHandler {
	return &PensonsHandler{
		service: personService,
	}
}

func (h *PensonsHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var input service.CreateUserInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if errors := ValidatePayload(input); errors != nil {
		WriteValidationError(w, http.StatusBadRequest, errors)
		return
	}

	output, err := h.service.CreateUser(r.Context(), input)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	WriteJSON(w, http.StatusCreated, output)
}

func (h *PensonsHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	personID := r.PathValue("personID")
	var input service.UpdateUserInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input.UserID = personID

	if errors := ValidatePayload(input); errors != nil {
		WriteValidationError(w, http.StatusBadRequest, errors)
		return
	}

	output, err := h.service.UpdateUser(r.Context(), input)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	WriteJSON(w, http.StatusCreated, output)
}

func (h *PensonsHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	personID := r.PathValue("personID")

	output, err := h.service.GetUser(r.Context(), personID)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	WriteJSON(w, http.StatusOK, output)
}

func (h *PensonsHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	personID := r.PathValue("personID")

	if err := h.service.DeleteUser(r.Context(), personID); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	WriteJSON(w, http.StatusNoContent, nil)
}
