package main

import (
	"encoding/json"
)

//Interface 
type atualizarEntrega interface {
	AtualizarStatus() 
}

//Struct da entrega
type entrega struct {
	Entregador string `json:"entregador"`
	Produto   string `json:"produto"`
	Quantidade int    `json:"quantidade"`
	Status []string `json:"status"`
}

//Método para atualizar o status da entrega

func main() {

	//lista de estatus 
	status := []string{"pendente", "em andamento", "entregue"}
}