package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Servidor HTTP Conectado")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
