package main

import (
	f "fmt"
	"math/rand"
)

// sem retorno e sem argumentos
func menuPrato (){
	f.Println(`Pratos da casa:
	0 - Nenhum prato;
	1 - Bife acebolado - R$20.0;
	2 - Arroz amanteigado e farofa - R$15.0;
	3 - Macarrão com queijo, presunto e brócolis - R$18.0;
	4 - Feijão preto com bacon e salada de batata - R$12.0;`)
}

func menuBebida (){
	f.Println(`Bebidas da casa:
	0 - Nenhuma bebida;
	1 - Refrigerante - R$5.0;
	2 - Suco - R$7.0;
	3 - Cerveja - R$10.0;`)
}

func numPedido () int {
	rand.Seed(rand.Int63()) //inicializa o gerador de números aleatórios
	return rand.Intn(500) //gera um número aleatório entre 0 e 499
}

// com retorno e argumentos 
func prato (numPrato int) string {
	switch numPrato {
	case 0:
		return "Nenhum prato"
	case 1:
		return "Bife acebolado"
	case 2:
		return "Arroz amanteigado"
	case 3:
		return "Macarrão com queijo, presunto e brócolis"
	case 4:
		return "Feijão preto com bacon e salada de batata"
	default:
		return "Prato não encontrado"
	}
}


// com retorno e argumentos 
func bebida (numPrato int) string {
	switch numPrato {
	case 0:
		return "Sem bebida"
	case 1:
		return "Refrigetante"
	case 2:
		return "Suco"
	case 3:
		return "Cerveja"
	default:
		return "Prato não encontrado"
	}
}

func valorPratoEBebida (numPrato, numBebida int) float64 {
	valores := map[string]map[int]float64{
		"Pratos": {
			0: 0.0,
			1: 20.0,
			2: 15.0,
			3: 18.0,
			4: 12.0,
		},
		"Bebidas": {
			0: 0.0,
			1: 5.0,
			2: 7.0,
			3: 10.0,
		},
	}

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


// retornando mais de um valor 
func notaPagamento (valor float64) (float64, float64) {
	taxa := 0.10 //10% de taxa de serviço
	taxaPedido := valor * taxa //calcula a taxa de serviço
	valorTaxa := valor + valor * taxa //adiciona a taxa ao valor total

	return taxaPedido, valorTaxa 
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

func mostrarStatus(pedidos []string) {
	if len(pedidos) == 0 {
		f.Println("Nenhum item no pedido.")
		return
	}
	for i, item := range pedidos {
		f.Printf("%d. %s\n", i+1, item)
	}
}
