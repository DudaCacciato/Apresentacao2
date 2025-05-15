package main

<<<<<<< HEAD
import "fmt"

=======
>>>>>>> 3b61749a05b6e0512b63749bf650dea3c0dfbfe9
func main() {

	//Refeições
	meal1 := Meal{
		Drink: []string{"Cerveja", "Refrigerante"},
		Food: []string{"Arroz amanteigado", "Macarrão com queijo, presunto e brócolis"},
	}

	meal2 := Meal{
		Drink: []string{"Suco", "Água"},
		Food: []string{"Bife acebolado"},
	}

	meal3 := Meal{
		Drink: []string{"Suco"},
		Food:  []string{"Feijão preto com bacon e salada de batata", "Arroz amanteigado"},
	}

	meal4 := Meal{
		Drink: []string{"Água", "Refrigerante"},
		Food:  []string{"Arroz amanteigado", "Bife acebolado"},
	}

	meal5 := Meal{
		Drink: []string{"Cerveja", "Suco"},
		Food:  []string{"Macarrão com queijo, presunto e brócolis", "Feijão preto com bacon e salada de batata"},
	}
<<<<<<< HEAD

	//Canal de pedidos
	order1 := Order(meal1, meal2)
	order2 := Order(meal3)
	order3 := Order(meal4, meal5)

	//Receitas
	receipt1 := &Receipt{}
	receipt2 := &Receipt{}
	receipt3 := &Receipt{}

	//Adicionando os pedidos a receita
	receipt1.AddOrder(order1)
	receipt2.AddOrder(order2)
	receipt3.AddOrder(order3)

	//Calculando o total
	receipt1.CalculateTotal()
	receipt2.CalculateTotal()
	receipt3.CalculateTotal()

	//Calculando a taxa
	receipt1.CalculateFeeAndFinal()
	receipt2.CalculateFeeAndFinal()
	receipt3.CalculateFeeAndFinal()

	//Convertendo para JSON
	json1, _ := receipt1.ToJSON()
	json2, _ := receipt2.ToJSON()
	json3, _ := receipt3.ToJSON()

	//Imprimindo o resultado
	for i, json := range [][]byte{json1, json2, json3} {
		fmt.Printf("Pedido %d: %s\n", i+1, json)
	}

}
=======
}
>>>>>>> 3b61749a05b6e0512b63749bf650dea3c0dfbfe9
