package modulos

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

////// Estruturas [INICIO]

// Utilizador Estrutura para a coleção "utilizadores"
type Utilizador struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Nome     string        `bson:"username"`
	Password string        `bson:"password"`
	Email    string        `bson:"email"`
}

// Tarefa Estrutura para a coleção "tarefas"
type Tarefa struct {
	ID                bson.ObjectId `bson:"_id,omitempty"`
	id                int           `bson:"id"`
	Titulo            string        `bson:"titulo"`
	Conteudo          string        `bson:"conteudo"`
	DataDeCriacao     time.Time     `bson:"data_de_criacao"`
	UltimaModificacao time.Time     `bson:"ultima_modificacao"`
	DataDeFim         time.Time     `bson:"data_de_fim"`
	Prioridade        int           `bson:"prioridade"`
	Categoria         int           `bson:"categoria"`
	Status            int           `bson:"status"`
	ExpiraEm          time.Time     `bson:"expira_em"`
	Utilizador        int           `bson:"utilizador"`
	Invisivel         int           `bson:"invisivel"`
}

////// Estruturas [FIM]

////// Funções de apoio [INICIO]

// obterColecao Retorna a sessão e a coleção com base na string dada
func obterColecao(nomeDaColecao string) (mgo.Session, mgo.Collection) {
	session, err := mgo.Dial("127.0.0.1") // Conecta-se à base de dados pelo IP
	if err != nil {                       // Controlo de erros
		panic(err)
	}

	// Conecta-se à coleção dada
	c := session.DB("Tasks").C(nomeDaColecao)

	return *session, *c
}

////// Funções de apoio [FIM]

////// Obter dados [INICIO]

// ObterUtilizador Obtem todos os utilizadores na coleção "utilizadores" e retorna-os
func ObterUtilizador() []Utilizador {
	session, c := obterColecao("utilizadores")

	defer session.Close() // Fecha a sessão à base de dados quando a função terminar (defer)

	var results []Utilizador // Slice que guarda variáveis do tipo Utilizador

	// Procura todos os utilizadores na coleção utilizadores
	// Ordena-os pela ordem de adição
	// Guarda-os no slice results
	err := c.Find(bson.M{}).Sort("-timestamp").All(&results)
	if err != nil { // Controlo de erros
		panic(err)
	}

	return results
}

// ObterTarefas Obtem todas as tarefas na coleção "tarefas" e retorna-os
func ObterTarefas(utilizadorID int) []Tarefa {
	session, c := obterColecao("tarefas")

	defer session.Close() // Defer

	var results []Tarefa // Slice que guarda variáveis do tipo Tarefas

	err := c.Find(bson.M{"utilizador": utilizadorID}).Sort("-timestamp").All(&results)
	if err != nil { // Controlo de erros
		panic(err)
	}

	return results
}

////// Obter dados [FIM]

////// Mudar dados (Update) [INICIO]

// MudarStatusTarefa Altera o estado de uma tarefa pelo definido pelo utilizador
func MudarStatusTarefa(utilizadorID int, tarefaID int, status int) bool {
	session, c := obterColecao("tarefas")

	defer session.Close()

	tarefaEspecifica := bson.M{"utilizador": utilizadorID, "id": tarefaID} // Tarefa a mudar
	alteracao := bson.M{"$set": bson.M{"status": status}}                  // Alterações a fazer
	err := c.Update(tarefaEspecifica, alteracao)                           // Escrita na Base de Dados

	if err != nil { // Controlo de erros
		return false
	}

	return true
}

// ReciclarTarefa Manda a tarefa para a lixeira ou elimina permanentemente a tarefa se ela já estiver lá
// string: Nome da tarefa; string: reciclada || eliminada; bool: Ocorreu erro
func ReciclarTarefa(utilizadorID int, tarefaID int) (string, string, bool) {
	session, c := obterColecao("tarefas")

	defer session.Close()

	tarefaEspecificaObj := Tarefa{}
	tarefaEspecificaBSON := bson.M{"utilizador": utilizadorID, "id": tarefaID}

	// Procura a tarefa específica
	err := c.Find(tarefaEspecificaBSON).One(&tarefaEspecificaObj)
	if err != nil {
		return "", "erro", false
	}

	// Elimina permanentemente caso já esteja na reciclagem
	// 2 equivale sempre a DELETED, ou seja, a tarefa já está na lixeira
	if tarefaEspecificaObj.Status == 2 {
		err = c.Remove(tarefaEspecificaBSON)
		// Se ocorrer algum erro
		if err != nil {
			return tarefaEspecificaObj.Titulo, "eliminar", true
		}
		// Se tudo correr bem
		return tarefaEspecificaObj.Titulo, "eliminada", false
	}

	// Move para a reciclagem
	// 2 equivale sempre a DELETED, ou seja, a tarefa vai para a lixeira
	err = c.Update(tarefaEspecificaBSON, bson.M{"$set": bson.M{"status": 2}})

	if err != nil {
		fmt.Println(err)
		return tarefaEspecificaObj.Titulo, "reciclar", true
	}
	return tarefaEspecificaObj.Titulo, "reciclada", false
}

////// Mudar dados (Update) [FIM]
