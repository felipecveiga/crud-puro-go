package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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

	body, _ := json.Marshal(user)
	r := bytes.NewReader(body)
	url := "/create"
	request := httptest.NewRequest("POST", url, r)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	handler.Create(response, request)

	if response.Code != http.StatusCreated {
		t.Errorf("erro no status code, esperado 201 e retornado foi %d", response.Code)
		return
	}
}
