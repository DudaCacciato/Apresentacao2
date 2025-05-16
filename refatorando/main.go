package main

import "fmt"

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

	// Adicionando as refeições 

	orders := []Meal{meal1, meal2, meal3, meal4, meal5}

	orderChannel := make(chan Meal)

	go func() {
		for _, order := range orders {
			orderChannel <- order
		}
		close(orderChannel)
	}()


	for meal := range orderChannel {
		receipt := &Receipt{}

		receipt.AddOrder(meal)
		receipt.CalculateTotal()
		receipt.CalculateFeeAndFinal()

		jsonData, err := receipt.ToJSON()
		if err != nil {
			fmt.Println("Erro ao converter para JSON:", err)
			return
		}
		fmt.Println(string(jsonData))
	}


}
