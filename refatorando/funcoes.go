package main

//Importações
import (
	"encoding/json"
	"math/rand"
)

// Interface para funções de pedido
type OrderInterface interface {
	Order(chan Meal)
}

// Interface para funções de nota fiscal
type ReceiptInterface interface {
	AddOrder(chan Meal)
	CalculateFeeAndFinal()
	CalculateTotal()
	ToJSON() ([]byte, error)
}

// Structs
type Meal struct {
	drink []string
	food  []string
}

type Receipt struct {
	ID 			 int    `json:"id"`
	OrderChannel []Meal  `json:"order"`
	FeeAmount    float64 `json:"fee_amount"`
	TotalAmount  float64 `json:"total_amount"`
	FinalAmount  float64 `json:"final_amount"`
}

//Funções

// Função que insere o pedido no canal de pedidos
func Order(meal ...Meal) <-chan Meal {

	c := make(chan Meal)

	go func() {
		for _, m := range meal {
			c <- m
		}
	}()

	return c
}

// Função que insere os dados do canal de pedidos a nota fiscal
func (r *Receipt) AddOrder(c <-chan Meal) {

	// Adiciona o ID da nota fiscal
	rand.Seed(rand.Int63())
	r.ID = rand.Intn(500)

	//Adicionando os pedidos a nota fiscal
	for m := range c {
		r.OrderChannel = append(r.OrderChannel, m)
	}
}

// Função que calcula a taxa de serviço e o valor final com taxa e adiciona na struct
func (r *Receipt) CalculateFeeAndFinal() {
	r.FeeAmount = 0.10 * r.TotalAmount
	r.FinalAmount = r.TotalAmount + r.FeeAmount
}

// Função que calcula o total sem taxa e adiciona na struct
func (r *Receipt) CalculateTotal() {}

// Função que converte a nota fiscal para json
func (r *Receipt) ToJSON() ([]byte, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// Função que irá receber os pedidos e retornar a ordem de pedidos de prioridade
func PriorityOrder(orders ...chan Meal) (chan Meal) {

	c := make(chan Meal)
	
	go func() {
		for _, orderChan := range orders {
            for order := range orderChan {
                c <- order
            }
        }
	}()

	return c
}

// Map de bebidas
func DrinksMenu() map[string]float64 {
	return map[string]float64{
		"Refrigerante": 5.0,
		"Suco":        7.0,
		"Cerveja":     10.0,
		"Água":       3.0,
	}
}

// Map de comidas
func FoodsMenu() map[string]float64 {
	return map[string]float64{
		"Bife acebolado": 20.0,
		"Arroz amanteigado": 15.0,
		"Macarrão com queijo, presunto e brócolis": 18.0,
		"Feijão preto com bacon e salada de batata": 12.0,
	}
}