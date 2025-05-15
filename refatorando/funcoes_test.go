package main

import (
	"testing"
)

const erro = "Valor esperado: %v, Valor recebido: %v"

func TestOrder(t *testing.T) {
	t.Parallel()

	meal1 := Meal{
		Drink: []string{"Cerveja", "Refrigerante"},
		Food: []string{"Arroz amanteigado", "Macarrão com queijo, presunto e brócolis"},
	}

	meal2 := Meal{
		Drink: []string{"Suco", "Água"},
		Food:  []string{"Bife acebolado"},
	}

	result := Order(meal1, meal2)

	if len(result) != 2 {
		t.Errorf(erro, 2, len(result))
	}
}

func TestAddOrder(t *testing.T) {
	t.Parallel()

	meal1 := Meal{
		Drink: []string{"Cerveja", "Refrigerante"},
		Food:  []string{"Arroz amanteigado", "Macarrão com queijo, presunto e brócolis"},
	}

	meal2 := Meal{
		Drink: []string{"Suco", "Água"},
		Food:  []string{"Bife acebolado"},
	}

	c := make(chan Meal)
	go func() {
		c <- meal1
		c <- meal2
		close(c)
	}()

	r := &Receipt{}
	r.AddOrder(c)

	if len(r.OrderChannel) != 2 {
		t.Errorf(erro, 2, len(r.OrderChannel))
	}
}

func TestCalculateTotal(t *testing.T) {
	t.Parallel()

	meal1 := Meal{
		Drink: []string{"Cerveja", "Refrigerante"},                        
		Food: []string{"Arroz amanteigado", "Macarrão com queijo, presunto e brócolis"},
		
	}

	meal2 := Meal{
		Drink: []string{"Suco", "Água"}, 
		Food: []string{"Bife acebolado"}, 
		
	}

	expectedTotal := 48.0 + 30.0 

	c := make(chan Meal)
	go func() {
		c <- meal1
		c <- meal2
		close(c)
	}()

	r := &Receipt{}
	r.AddOrder(c)
	r.CalculateTotal()

	if r.TotalAmount != expectedTotal {
		t.Errorf("esperado %.2f, obtido %.2f", expectedTotal, r.TotalAmount)
	}
}
