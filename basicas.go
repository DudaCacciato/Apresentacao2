package main

import (
	f "fmt"
	"math/rand"
)


func valorPratoEBebida (numPrato, numBebida int) float64 {

	for prato, valorPrato := range valores["Pratos"] { //for percorrer o mapa de pratos
		if prato == numPrato { //se o prato for igual ao número do prato
			for bebida, valorBebida := range valores["Bebidas"] { //for percorrer o mapa de bebidas
				if bebida == numBebida { //se a bebida for igual ao número da bebida
					return valorPrato + valorBebida //retorna o valor do prato + o valor da bebida
				}
			}
		}
	}

	return 0.0 //se não encontrar o prato ou a bebida, retorna 0.0
}


func deletarItemPedido(itemDeletar int, pedidos []string) []string {
	if itemDeletar > 0 && itemDeletar <= len(pedidos) {
		index := itemDeletar - 1
		pedidos = append(pedidos[:index], pedidos[index+1:]...)
		f.Println("Item deletado com sucesso.")
	} else {
		f.Println("Item inválido.")
	}
	return pedidos
}
