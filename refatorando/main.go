package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

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

	// for _, meal := range orders {
	// 	go func(m Meal) {
	// 		AddChannel(m, orderChannel)
	// 		time.Sleep(1 * time.Second)
	// 	}(meal)
	// }

	wg := &sync.WaitGroup{}
	wg.Add(len(orders))

	for _, meal := range orders {
		go func(Meal) {
			rand.Seed(rand.Int63())
			nRandom := rand.Intn(5)
			time.Sleep(time.Duration(nRandom) * time.Millisecond)
			orderChannel <- meal
			wg.Done()
		}(meal)
	}

	wg.Wait()
	close(orderChannel)

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
