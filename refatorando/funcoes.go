package main

//Importações
import (
	"encoding/json"
	"math"
	"math/rand"
)

// Interface para funções de nota fiscal
type ReceiptInterface interface {
	AddOrder(chan Meal)
	CalculateFeeAndFinal()
	CalculateTotal()
	ToJSON() ([]byte, error)
}

// Structs
type Meal struct {
	Drink []string
	Food  []string
}

type Receipt struct {
	ID          int     `json:"id"`
	Order       []Meal  `json:"order"`
	FeeAmount   float64 `json:"fee_amount"`
	TotalAmount float64 `json:"total_amount"`
	FinalAmount float64 `json:"final_amount"`
}

//Funções

// Função que insere os dados do canal de pedidos a nota fiscal
func (r *Receipt) AddOrder(meal Meal) {

	// Adiciona o ID da nota fiscal
	rand.Seed(rand.Int63())
	r.ID = rand.Intn(500)

	//Adicionando os pedidos a nota fiscal
	r.Order = append(r.Order, meal)
}

// Função que calcula o total sem taxa e adiciona na struct
func (r *Receipt) CalculateTotal() {
	total := 0.0
	drinksMenu := DrinksMenu()
	foodsMenu := FoodsMenu()

	for _, m := range r.Order {
		for _, drink := range m.Drink {
			for name, price := range drinksMenu {
				if drink == name {
					total += price
				}
			}
		}
		for _, food := range m.Food {
			for name, price := range foodsMenu {
				if food == name {
					total += price
				}
			}
		}
	}

	r.TotalAmount = math.Round(total)
}

// Função que calcula a taxa de serviço e o valor final com taxa e adiciona na struct
func (r *Receipt) CalculateFeeAndFinal() {
	r.FeeAmount = math.Round(0.10 * r.TotalAmount)
	r.FinalAmount = math.Round(r.TotalAmount + r.FeeAmount)
}

// Função que converte a nota fiscal para json
func (r *Receipt) ToJSON() ([]byte, error) {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// Map de bebidas
func DrinksMenu() map[string]float64 {
	return map[string]float64{
		"Refrigerante": 5.0,
		"Suco":         7.0,
		"Cerveja":      10.0,
		"Água":         3.0,
	}
}

// Map de comidas
func FoodsMenu() map[string]float64 {
	return map[string]float64{
		"Bife acebolado":    20.0,
		"Arroz amanteigado": 15.0,
		"Macarrão com queijo, presunto e brócolis":  18.0,
		"Feijão preto com bacon e salada de batata": 12.0,
	}
}
