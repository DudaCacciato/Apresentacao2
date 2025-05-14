package main

func main() {

	//Pedidos

	meal1 := Meal{
		drink: []string{"Cerveja", "Refrigerante"},
		food:  []string{"Pizza", "Macarrão com queijo, presunto e brócolis"},
	}

	meal2 := Meal{
		drink: []string{"Suco", "Água"},
		food:  []string{"Bife acebolado"},
	}

	meal3 := Meal{
		drink: []string{"Suco"},
		food:  []string{"Feijão preto com bacon e salada de batata", "Arroz amanteigado"},
	}

	meal4 := Meal{
		drink: []string{"Água", "Refrigerante"},
		food:  []string{"Arroz amanteigado", "Bife acebolado"},
	}

	meal5 := Meal{
		drink: []string{"Cerveja", "Suco"},
		food:  []string{"Macarrão com queijo, presunto e brócolis", "Feijão preto com bacon e salada de batata"},
	}
}