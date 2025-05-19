package main

import "fmt"

func main() {
	//Refeições
	orders := []Meal{
		{
			Drink: []string{"Cerveja", "Refrigerante"},
			Food:  []string{"Arroz amanteigado", "Macarrão com queijo, presunto e brócolis"},
		},

		{
			Drink: []string{"Suco", "Água"},
			Food:  []string{"Bife acebolado"},
		},

		{
			Drink: []string{"Suco"},
			Food:  []string{"Feijão preto com bacon e salada de batata", "Arroz amanteigado"},
		},

		{
			Drink: []string{"Água", "Refrigerante"},
			Food:  []string{"Arroz amanteigado", "Bife acebolado"},
		},

		{
			Drink: []string{"Cerveja", "Suco"},
			Food:  []string{"Macarrão com queijo, presunto e brócolis", "Feijão preto com bacon e salada de batata"},
		},
	}
	// Adicionando as refeições

	orderChannel := make(chan Meal, len(orders))
	defer close(orderChannel)

	// for _, meal := range orders {
	// 	go func(m Meal) {
	// 		AddChannel(m, orderChannel)
	// 		time.Sleep(1 * time.Second)
	// 	}(meal)
	// }

	for _, meal := range orders {
		go func(Meal) {
			orderChannel <- meal
		}(meal)
	}

	// go func(){
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
	// }()

}
