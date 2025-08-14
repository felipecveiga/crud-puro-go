package handler

import (
	"encoding/json"
	"net/http"

	"github.com/felipecveiga/crud-puro-go/model"
	"github.com/felipecveiga/crud-puro-go/service"
)

//go:generate mockgen -source=./user.go -destination=./user_mock.go -package=handler
type Handler interface{
	Create(response http.ResponseWriter, request *http.Request)
}

type handler struct {
	Service service.Service
}

func NewUserHandler(s service.Service) Handler {
	return &handler{
		Service: s,
	}
}

func (h *handler) Create(response http.ResponseWriter, request *http.Request) {
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

	response.WriteHeader(http.StatusCreated)
}
