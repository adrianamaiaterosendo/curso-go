package main

import "fmt"

func main() {

	aprovados := make(map[int]string)

	aprovados[12345] = "Maria"
	aprovados[17366] = "Joao"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s {CPF: %d }\n", nome, cpf)
	}

	delete(aprovados, 12345)

	fmt.Println("Tem algum aprovado?", aprovados)

	funcionarioSalario := map[string]map[string]float64{
		"A": {
			"Adriana": 35000,
			"Antonio": 54666,
		},
		"B": {
			"Bruno": 20999,
		},
	}

	for letra, funcs := range funcionarioSalario {
		fmt.Println(funcs)
		fmt.Println(letra)
	}
}
