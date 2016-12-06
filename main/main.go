package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"fmt"

	db "github.com/Tasks/modulos"
)

func main() {
	// Relativo a tarefas específicas
	http.HandleFunc("/status/", AlterarStatusTarefa)
	http.HandleFunc("/eliminar/", EliminarTarefa)
	http.HandleFunc("/editar/", EditarTarefa)
	http.HandleFunc("/restaurar/", RestaurarTarefa)
	http.HandleFunc("/criar/", CriarTarefa)
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
	if erro { // Houve um erro
		message = "Ocorreu um erro a " + destino + " a tarefa " + strings.ToUpper(nome) + ". Tente novamente mais tarde"
	} else if destino == "erro" { // A tarefa não foi encontrada
		message = "A tarefa não existe!"
	} else { // A operação correu como esperado
		message = "A Tarefa " + strings.ToUpper(nome) + " foi " + destino + " com sucesso!"
	}

	w.Write([]byte(message))
}

// EditarTarefa Edita o conteúdo de uma tarefa com base nos parametros recebidos
func EditarTarefa(w http.ResponseWriter, r *http.Request) {
	tarefaModelo := db.Tarefa{}
	tarefaModelo.ID = 0
	tarefaModelo.Titulo = "Tarefa Alterada"
	tarefaModelo.Conteudo = "Esta tarefa foi alterada pelo utilizador"
	tarefaModelo.DataDeFim = time.Now()
	tarefaModelo.Prioridade = 2
	tarefaModelo.Categoria = 2
	tarefaModelo.Status = 1
	tarefaModelo.ExpiraEm = time.Unix(0, 0) // Mesmo que nil
	tarefaModelo.Utilizador = 1
	tarefaModelo.Invisivel = 0

	resultado := db.EditarTarefa(tarefaModelo)

	var message string

	if !resultado {
		message = "Ocorreu um erro a editar " + strings.ToUpper(tarefaModelo.Titulo)
	} else {
		message = "A tarefa " + strings.ToUpper(tarefaModelo.Titulo) + " foi editada com sucesso!"
	}

	w.Write([]byte(message))
}

// RestaurarTarefa Remove a tarefa da lixeira e reverte para o seu status anterior
func RestaurarTarefa(w http.ResponseWriter, r *http.Request) {
	resultado := db.RestaurarTarefa(0)

	var message string

	if !resultado {
		message = "Ocorreu um erro a tirar " + "<nome da tarefa>" + " da reciclagem"
	} else {
		message = "A tarefa " + "<nome da tarefa>" + " foi removida da reciclagem com sucesso!"
	}

	w.Write([]byte(message))
}

// CriarTarefa Cria uma nova tarefa
func CriarTarefa(w http.ResponseWriter, r *http.Request) {
	tarefaModelo := db.Tarefa{}
	tarefaModelo.Titulo = "Nova Tarefa"
	tarefaModelo.Conteudo = "Esta tarefa foi criada pelo utilizador"
	tarefaModelo.DataDeFim = time.Now()
	tarefaModelo.Prioridade = 2
	tarefaModelo.Categoria = 2
	tarefaModelo.Status = 1
	tarefaModelo.ExpiraEm = time.Unix(0, 0) // Mesmo que nil
	tarefaModelo.Utilizador = 1
	tarefaModelo.Invisivel = 0

	resultado := db.CriarTarefa(tarefaModelo)

	var message string

	if !resultado {
		message = "Não foi possível criar a tarefa " + strings.ToUpper(tarefaModelo.Titulo)
	} else {
		message = "A tarefa " + strings.ToUpper(tarefaModelo.Titulo) + " foi criada com sucesso!"
	}

	w.Write([]byte(message))
}
