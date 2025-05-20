package main

// Importações
import (
	"testing"
)

// Variável padrão para o retorno do teste 
const erro = "Valor esperado: %v, Valor recebido: %v"

// Teste da função AddOrder
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

	
	receipt := &Receipt{}

	for meal := range c {
		receipt.AddOrder(meal)
	}

	if len(receipt.Order) != 2 {
		t.Errorf(erro, 2, len(receipt.Order))
	}
}

// Teste da função CalculateTotal
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

	
	receipt := &Receipt{}

	for meal := range c {
		receipt.AddOrder(meal)
		receipt.CalculateTotal()
	}

	if receipt.TotalAmount != expectedTotal {
		t.Errorf("esperado %.2f, obtido %.2f", expectedTotal, receipt.TotalAmount)
	}
}

//Teste da função CalculateFeeAndFinal
func TestCalculateFeeAndFinal(t *testing.T) {
	t.Parallel()

	meal1 := Meal{
		Drink: []string{"Cerveja", "Refrigerante"},
		Food:  []string{"Arroz amanteigado", "Macarrão com queijo, presunto e brócolis"},
	}

	meal2 := Meal{
		Drink: []string{"Suco", "Água"},
		Food:  []string{"Bife acebolado"},
	}

	expectedFee := 8.0
	expectedFinal := 86.0

	c := make(chan Meal)
	go func() {
		c <- meal1
		c <- meal2
		close(c)
	}()

	
	receipt := &Receipt{}

	for meal := range c {
		receipt.AddOrder(meal)
		receipt.CalculateTotal()
		receipt.CalculateFeeAndFinal()
	}

	if receipt.FeeAmount != expectedFee {
		t.Errorf(erro, expectedFee, receipt.FeeAmount)
	}

	if receipt.FinalAmount != expectedFinal {
		t.Errorf(erro, expectedFinal, receipt.FinalAmount)
	}
}