package main

import (
	"log"
	"net/http"
)

func main() {
	// URLs
	http.HandleFunc("/", CompleteTaskFunc)

	log.Fatal(http.ListenAndServe(":8080", nil)) // O servidor está ativo em: localhost:8080
}

func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) { // "/"
	// Escreve para o corpo do documento
	w.Write([]byte("<h1>Bem-vindo ao meu site!</h1>"))
}
