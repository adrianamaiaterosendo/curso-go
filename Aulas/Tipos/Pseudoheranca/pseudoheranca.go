package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       //campos anonimos
	turboLigado bool
}

func main() {
	f1 := ferrari{}

	f1.nome = "Flex Turbo"
	f1.velocidadeAtual = 160
	f1.turboLigado = true

	fmt.Println(f1)
}
