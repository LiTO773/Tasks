package main

import (
	"log"
	"net/http"
)

func main() {
	PORT := "127.0.0.1:8080" // localhost:8080
	log.Print("O servidor está ativo em: " + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
