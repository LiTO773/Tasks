package main

import (
	"fmt"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

// Utilizador Estrutura para a coleção "utilizadores"
type Utilizador struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Nome     string
	Password string
	Email    string
}

// MostrarUtilizador Obtem todos os utilizadores na coleção "utilizadores"
func MostrarUtilizador(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("127.0.0.1") // Conecta-se à base de dados pelo IP`
	if err != nil {                       // Controlo de erros
		panic(err)
	}

	defer session.Close() // Fecha a sessão à base de dados quando a função terminar (defer)`

	// Conectar à coleção utilizadores
	c := session.DB("Tasks").C("utilizadores")

	var results []Utilizador // Slice que guarda variáveis do tipo Utilizador

	// Procura todos os utilizadores na coleção utilizadores
	// Ordena-os pela ordem de adição
	// Guarda-os no slice results
	err = c.Find(bson.M{}).Sort("-timestamp").All(&results)
	if err != nil { // Controlo de erros
		panic(err)
	}

	fmt.Println("Results All: ", results)
	fmt.Println("Função rodou!")
}
