package main

import (
	f "fmt"
	"strings"
)

func main() {
	// Variáveis
	var numPrato, numBebida int
	var pedidos []string
	var valorTotal float64
	var numeroPedido = numPedido()

	f.Println("\n------ SEJA BEM-VINDO! ------")
	f.Println("\n O número do pedido é: ", numeroPedido)

	for {
		f.Println("\n------ MENU ------")
		menuPrato()
		menuBebida()

		for{
			// Escolha do prato
			f.Print("\nEscolha o prato: ")
			_, err1 := f.Scan(&numPrato)
			if err1 != nil {
				f.Println("Erro ao ler a entrada:", err1)
				continue
			}

			if numPrato == 0 {
				f.Println("Nenhum prato escolhido.")
			} else {
				f.Println("Prato do pedido: " + prato(numPrato))
				pedidos = append(pedidos, prato(numPrato))
				break
			}
		}

		for{
			// Escolha da bebida
			f.Print("Escolha a bebida: ")
			_, err2 := f.Scan(&numBebida)
			if err2 != nil {
				f.Println("Erro ao ler a entrada:", err2)
				continue
			}

			if numBebida == 0 {
				f.Println("Nenhuma bebida escolhida.")
				break
			} else {
				f.Println("Bebida do pedido: " + bebida(numBebida))
				pedidos = append(pedidos, bebida(numBebida))
				break
			}
		}

		valorTotal += valorPratoEBebida(numPrato, numBebida)

		f.Println("\n------ STATUS DO PEDIDO ------")
		mostrarStatus(pedidos)

		var resposta string

		for {
			f.Println("\nDeseja fazer outro pedido? (s/n)")

			_, err3 := f.Scan(&resposta)
			if err3 != nil {
				f.Println("Erro ao ler a entrada:", err3)
				continue
			}

			resposta = strings.ToLower(resposta)
	
			if resposta == "s" || resposta == "n" {
				break
			}

			f.Println("Opção inválida.")
			continue
		}

		if resposta == "n" {
			for {
				f.Println("Deseja deletar algum item do pedido atual? (s/n)")
				var rDeletar string
				_, err4 := f.Scan(&rDeletar)
				if err4 != nil {
					f.Println("Erro ao ler a entrada:", err4)
					continue
				}

				if rDeletar == "n" {
					break
				}

				for {
					f.Println("Digite o número do item que deseja deletar:")
					var itemDeletar int
					_, err := f.Scan(&itemDeletar)
					if err != nil {
						f.Println("Entrada inválida. Digite um número inteiro válido.")
						f.Scanln()
						continue
					}
				
					if itemDeletar < 1 || itemDeletar > len(pedidos) && len(pedidos) != 0 {
						f.Printf("Número inválido. Digite um número entre 1 e %d.\n", len(pedidos))
						continue
					}

					if len(pedidos) == 0 {
						f.Println("Nenhum item no pedido.")
						break
					}
				
					pedidos = deletarItemPedido(itemDeletar, pedidos)
					f.Println("\nPedido atualizado!")
					f.Println("\n------ STATUS DO PEDIDO ------")
					mostrarStatus(pedidos)
					break
				}

				if len(pedidos) == 0 {
					break
				}
			}
		}
		break
	}

	// ---------- NOTA FISCAL ----------

	f.Println("\n ****** Nota Fiscal ******")
	for _, pedido := range pedidos {
		f.Println(pedido)
	}

	f.Println("\n- Número do pedido: ", numeroPedido)

	taxa, valorFinal := notaPagamento(valorTotal)
	f.Printf("- Valor total sem taxa: R$%.2f", valorTotal)
	f.Printf("\n- Taxa de serviço: R$%.2f", taxa)
	f.Printf("\n- Valor total com taxa: R$%.2f", valorFinal)

	f.Println("\n\n AGRADECEMOS A PREFERÊNCIA! 🤗")
	f.Println("\n *********************")
}
