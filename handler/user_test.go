package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/felipecveiga/crud-puro-go/errs"
	"github.com/felipecveiga/crud-puro-go/model"
	"github.com/felipecveiga/crud-puro-go/service"
	"go.uber.org/mock/gomock"
)

func TestCreateUserHandler_WhenReturSucess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	handler := NewUserHandler(mockService)

	user := model.User{
		Name:  "Felipe",
		Email: "felipe@gmail.com",
		Sexo:  "Masculino",
		Age:   31,
		Phone: 212121,
		Residence: model.Residence{
			Street:  "Brasil",
			City:    "rio de janeiro",
			Country: "rua a",
			Number:  27,
		},
	}

	mockService.EXPECT().
		CreateUser(gomock.Any()).
		Return(nil)

	endpoint := "/create"
	body, _ := json.Marshal(user)
	request := httptest.NewRequest("POST", endpoint, bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	handler.Create(response, request)

	if response.Code != http.StatusCreated {
		t.Errorf("erro no status code, esperado: %d, retornado: %d", http.StatusCreated, response.Code)
		return
	}
}

func TestCreateUserHandler_WhenReturError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	handler := NewUserHandler(mockService)

	user := model.User{
		Name:  "Felipe",
		Email: "felipe@gmail.com",
		Sexo:  "Masculino",
		Age:   31,
		Phone: 21212121,
		Residence: model.Residence{
			Street:  "Brasil",
			City:    "Rio de Janeiro",
			Country: "Rua ABC",
			Number:  30,
		},
	}

	endpoint := "/create"
	body, _ := json.Marshal(user)
	request := httptest.NewRequest("POST", endpoint, bytes.NewReader(body))
	response := httptest.NewRecorder()
	request.Header.Set("Content-Type", "application/json")

	mockService.EXPECT().
		CreateUser(gomock.Any()).
		Return(errs.ErrBodyRequest)

	handler.Create(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("erro no status code, esperado: %d, retornado: %d", http.StatusBadRequest, response.Code)
		return
	}
}

func TestCreateUserHandler_WhenReturErrorMethodRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	handler := NewUserHandler(mockService)

	endpoint := "/create"
	request := httptest.NewRequest("GET", endpoint, nil)
	response := httptest.NewRecorder()
	request.Header.Set("Content-Type", "application/json")

	handler.Create(response, request)

	if response.Code != http.StatusMethodNotAllowed {
		t.Errorf("erro no método da requisição, erro retornado: %d", response.Code)
		return
	}
}

func TestCreateUserHandler_WhenReturErrorBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	handler := NewUserHandler(mockService)

	bodyInvalido := []byte(`{nome: "felipe"}`)

	endpoint := "/create"
	response := httptest.NewRecorder()
	request := httptest.NewRequest("POST", endpoint, bytes.NewBuffer(bodyInvalido))
	request.Header.Set("Content-Type", "application/json")

	handler.Create(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Status code esperado %d, retornado %d", http.StatusBadRequest, response.Code)
	}
}

func TestGetUserHandler_WhenReturnSucess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	handler := NewUserHandler(mockService)

	endpoint := "/user/68a8e66a5a3b238655f42f4"
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", endpoint, nil)

	user := model.User{
		Name:  "Felipe",
		Email: "felipe@gmail.com",
		Sexo:  "Masculino",
		Age:   31,
		Phone: 21212121,
		Residence: model.Residence{
			Street:  "Brasil",
			City:    "Rio de Janeiro",
			Country: "Rua ABC",
			Number:  30,
		},
	}

	mockService.EXPECT().
		GetUser("68a8e66a5a3b238655f42f4").
		Return(&user, nil)

	handler.GetUser(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Status code esperado %d, retornado %d", http.StatusOK, response.Code)
	}
}

func TestGetUserHandler_WhenReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	handler := NewUserHandler(mockService)

	endpoint := "/user/68a8e66a5a3b238655f42f4"
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", endpoint, nil)

	mockService.EXPECT().
		GetUser("68a8e66a5a3b238655f42f4").
		Return(nil, errors.New("some error"))

	handler.GetUser(response, request)
	if response.Code != http.StatusBadRequest {
		t.Errorf("Status code esperado %d, retornado %d", http.StatusBadRequest, response.Code)
	}
}

func TestGetUserHandler_WhenReturErrorMethodRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	handler := NewUserHandler(mockService)

	endpoint := "/user/68a8e66a5a3b238655f42f4"
	request := httptest.NewRequest("POST", endpoint, nil)
	response := httptest.NewRecorder()

	handler.GetUser(response, request)

	if response.Code != http.StatusMethodNotAllowed {
		t.Errorf("erro no método da requisição, erro retornado: %d", response.Code)
		return
	}
}

func TestGetUserHandler_WhenReturErrorEndPoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	handler := NewUserHandler(mockService)

	endpoint := "/use/68a8e66a5a3b238655f42f4"
	request := httptest.NewRequest("GET", endpoint, nil)
	response := httptest.NewRecorder()

	handler.GetUser(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Status code esperado %d, retornado %d", http.StatusBadRequest, response.Code)
		return
	}
}

func TestGetUserHandler_WhenReturnErrUserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	handler := NewUserHandler(mockService)

	endpoint := "/user/68a8e66a5a3b238655f42f4"
	request := httptest.NewRequest("GET", endpoint, nil)
	response := httptest.NewRecorder()

	mockService.EXPECT().
		GetUser("68a8e66a5a3b238655f42f4").
		Return(nil, errs.ErrUserNotFound)

	handler.GetUser(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Status code esperado %d, retornado %d", http.StatusNotFound, response.Code)
		return
	}
}
