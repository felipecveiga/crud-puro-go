package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/felipecveiga/crud-puro-go/config/db"
	"github.com/felipecveiga/crud-puro-go/handler"
	"github.com/felipecveiga/crud-puro-go/repository"
	"github.com/felipecveiga/crud-puro-go/service"
)

func main() {

	clientDB := db.Connection()

	repository := repository.NewUserRepository(clientDB)
	service := service.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	http.HandleFunc("/create", handler.Create)
	http.HandleFunc("/user/", handler.GetUser)
	http.HandleFunc("/users/", handler.GetUser)

	fmt.Println("Servidor HTTP Conectado")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
