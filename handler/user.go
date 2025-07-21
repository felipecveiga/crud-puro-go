package handler

import (
	"encoding/json"
	"net/http"

	"github.com/felipecveiga/crud-puro-go/model"
	"github.com/felipecveiga/crud-puro-go/service"
)

func Create(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "método HTTP inválido para requisição", http.StatusMethodNotAllowed)
	}

	payload := new(model.User)
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		http.Error(response, "erro no body da requisição", http.StatusBadRequest)
	}

	result, err := service.CreateUser(payload)
	if err != nil {
		http.Error(response, "erro ao criar conta do usuário", http.StatusBadRequest)
	}
	
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
