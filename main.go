package main

import (
	"log"
	"net/http"
)

func main() {
	PORT := "127.0.0.1:8080" // localhost:8080
	log.Print("O servidor está ativo em: " + PORT)

	// URLs
	http.HandleFunc("/", CompleteTaskFunc)

	log.Fatal(http.ListenAndServe(PORT, nil))
}

func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) { // "/"
	// Escreve para o corpo do documento
	w.Write([]byte("<h1>Bem-vindo ao meu site!</h1>"))
}
