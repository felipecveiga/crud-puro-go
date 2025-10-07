package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/felipecveiga/crud-puro-go/errs"
	"github.com/felipecveiga/crud-puro-go/model"
	"github.com/felipecveiga/crud-puro-go/service"
)

//go:generate mockgen -source=./user.go -destination=./user_mock.go -package=handler
type Handler interface {
	Create(response http.ResponseWriter, request *http.Request)
	GetUser(response http.ResponseWriter, request *http.Request)
	GetAllUsers(response http.ResponseWriter, request *http.Request)
	DeleteUser(response http.ResponseWriter, request *http.Request)
	UpdateUser(response http.ResponseWriter, request *http.Request)
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
		http.Error(response, errs.ErrInvalidHTTPMethod.Error(), http.StatusMethodNotAllowed)
		return
	}

	payload := new(model.User)
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		http.Error(response, errs.ErrInvalidPayload.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.CreateUser(payload)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

func (h *handler) GetUser(response http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodGet {
		http.Error(response, errs.ErrInvalidHTTPMethod.Error(), http.StatusMethodNotAllowed)
		return
	}

	endpoint := request.URL.Path
	separadorUrl := strings.Split(endpoint, "/")
	var id string

	if len(separadorUrl) >= 3 && separadorUrl[1] == "user" {
		id = strings.TrimSpace(separadorUrl[2])
	} else {
		http.Error(response, errs.ErrUserID.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetUser(id)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			http.Error(response, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}

func (h *handler) GetAllUsers(response http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodGet {
		http.Error(response, errs.ErrInvalidHTTPMethod.Error(), http.StatusMethodNotAllowed)
		return
	}

	users, err := h.Service.GetAllUsers()
	if err != nil {
		if errors.Is(err, errs.ErrUsersNotFound) {
			http.Error(response, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}

func (h *handler) DeleteUser(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodDelete {
		http.Error(response, errs.ErrInvalidHTTPMethod.Error(), http.StatusMethodNotAllowed)
		return
	}

	endpoint := request.URL.Path
	separadorURL := strings.Split(endpoint, "/")

	var id string
	if len(separadorURL) >= 3 && separadorURL[1] == "delete" {
		id = strings.TrimSpace(separadorURL[2])
	} else {
		http.Error(response, errs.ErrUserID.Error(), http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteUser(id)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			http.Error(response, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}

func(h *handler) UpdateUser(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPatch {
		http.Error(response, errs.ErrInvalidHTTPMethod.Error(), http.StatusMethodNotAllowed)
		return
	}

	endpoint := request.URL.Path
	separadorURL := strings.Split(endpoint, "/")
	
	var id string
	if len(separadorURL) >= 3 && separadorURL[1] == "update" {
		id = strings.TrimSpace(separadorURL[2])
	} else {
		http.Error(response, errs.ErrIDInvalid.Error(), http.StatusBadRequest)
		return
	}

	payload := new(model.User)
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		http.Error(response, errs.ErrInvalidPayload.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.UpdateUser(id, payload)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			http.Error(response, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}
