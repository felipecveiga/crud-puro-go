package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/felipecveiga/crud-puro-go/handler"
)

func main() {

	http.HandleFunc("/create", handler.Create)

	fmt.Println("Servidor HTTP Conectado")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
