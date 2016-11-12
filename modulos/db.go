package modulos

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Utilizador Estrutura para a coleção "utilizadores"
type Utilizador struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Nome     string
	Password string
	Email    string
}

// ObterUtilizador Obtem todos os utilizadores na coleção "utilizadores" e retorna-os
func ObterUtilizador() []Utilizador {
	session, err := mgo.Dial("127.0.0.1") // Conecta-se à base de dados pelo IP`
	if err != nil {                       // Controlo de erros
		panic(err)
	}

	defer session.Close() // Fecha a sessão à base de dados quando a função terminar (defer)

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

	return results
}
