package main

import (
	"log"
	"net/http"

	"fmt"

	db "github.com/Tasks/modulos"
)

func main() {
	// Relativo a tarefas específicas
	// http.HandleFunc("/terminar/", TerminarTarefa)
	// http.HandleFunc("/eliminar/", EliminarTarefa)
	// http.HandleFunc("/editar/", EditarTarefa)
	// http.HandleFunc("/restaurar/", RestaurarTarefa)
	// http.HandleFunc("/criar/", CriarTarefa)
	// http.HandleFunc("/atualizar/", AtualizarTarefa)
	// http.HandleFunc("/procurar/", ProcurarTarefa)

	// // Relativo a todas as tarefas
	http.HandleFunc("/", MostrarUtilizador)
	// http.HandleFunc("/eliminado/", TarefasEliminadas)
	// http.HandleFunc("/reciclagem/", TarefasRecicladas)
	// http.HandleFunc("/terminado/", TarefasTerminadas)

	// // Relativo ao utilizador
	// http.HandleFunc("/entrar", Login)
	// http.HandleFunc("/registar", Signin)
	// http.HandleFunc("/admin", Administracao)
	// http.HandleFunc("/adicionar_utilizador", AdicionarUtilizador)
	// http.HandleFunc("/alterar", AlterarUtilizador)
	// http.HandleFunc("/sair", Logout)

	// Ficheiros estáticos
	http.Handle("/static/", http.FileServer(http.Dir("../public")))

	log.Print("O servidor está ativo na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// MostrarTarefas Mostrar todas tarefas pendentes do utilizador ("/")
func MostrarTarefas(w http.ResponseWriter, r *http.Request) {
	message := "Todas as tarefas pendentes (" + r.Method + ")" //r.Method == "GET" || r.Method == "POST"
	w.Write([]byte(message))
}

// MostrarUtilizador Mostrar informação do utilizador
func MostrarUtilizador(w http.ResponseWriter, r *http.Request) {
	message := "Todas as tarefas pendentes (" + r.Method + ")" //r.Method == "GET" || r.Method == "POST"
	utilizador := db.ObterUtilizador()
	fmt.Println(utilizador)
	w.Write([]byte(message))
}
