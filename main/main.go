package main

import (
	"log"
	"net/http"
	"strings"

	"fmt"

	db "github.com/Tasks/modulos"
)

func main() {
	// Relativo a tarefas específicas
	http.HandleFunc("/status/", AlterarStatusTarefa)
	http.HandleFunc("/eliminar/", EliminarTarefa)
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
	tarefas := db.ObterTarefas(1)
	fmt.Println(tarefas)
	w.Write([]byte(message))
}

// MostrarUtilizador Mostrar informação do utilizador
func MostrarUtilizador(w http.ResponseWriter, r *http.Request) {
	message := "Todas as tarefas pendentes (" + r.Method + ")"
	utilizador := db.ObterUtilizador()
	fmt.Println(utilizador)
	w.Write([]byte(message))
}

// AlterarStatusTarefa Altera o status da tarefa com base no indicado pelo utilizador
func AlterarStatusTarefa(w http.ResponseWriter, r *http.Request) {
	message := "Todas as tarefas pendentes (" + r.Method + ")"
	err := db.MudarStatusTarefa(1, 0, 3) // Teste (funciona)
	fmt.Println(err)
	w.Write([]byte(message))
}

// EliminarTarefa Elimina a tarefa ou move-a para a reciclagem
func EliminarTarefa(w http.ResponseWriter, r *http.Request) {
	nome, destino, erro := db.ReciclarTarefa(1, 1)
	var message string
	if erro { // Caso haja um erro
		message = "Ocorreu um erro a " + destino + " a tarefa " + strings.ToUpper(nome) + ". Tente novamente mais tarde"
	} else if destino == "erro" { // A tarefa não foi encontrada
		message = "A tarefa não existe!"
	} else { // A operação ocorreu como esperado
		message = "A Tarefa " + strings.ToUpper(nome) + " foi " + destino + " com sucesso!"
	}
	w.Write([]byte(message))
}

// {
//     "_id" : ObjectId("58190caee55240d7ac1cf139"),
//     "id" : 1,
//     "titulo" : "gofmtall",
//     "conteudo" : "The idea is to run go fmt -w file.go on every go file in the listing, *Edit turns out this difficult to do in golang **Edit brely 3 line bash script. ",
//     "data_de_criacao" : ISODate("2015-11-12T16:58:31.000Z"),
//     "ultima_modificacao" : ISODate("2015-11-14T10:42:14.000Z"),
//     "data_de_fim" : ISODate("2015-11-13T13:16:48.000Z"),
//     "prioridade" : 3,
//     "categoria" : 1,
//     "status" : 1,
//     "expira_em" : null,
//     "utilizador" : 1,
//     "invisivel" : 0
// }
