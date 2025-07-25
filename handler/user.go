package handler

import (
	"encoding/json"
	"net/http"

	"github.com/felipecveiga/crud-puro-go/model"
	"github.com/felipecveiga/crud-puro-go/service"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) Create(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "método HTTP inválido para requisição", http.StatusMethodNotAllowed)
	}

	payload := new(model.User)
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		http.Error(response, "erro no body da requisição", http.StatusBadRequest)
	}

	err = h.Service.CreateUser(payload)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusOK)
}
