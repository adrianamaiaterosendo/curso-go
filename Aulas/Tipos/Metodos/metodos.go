package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomecompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")

	p.nome = partes[0]
	p.sobrenome = partes[1]

}

func main() {
	p1 := pessoa{
		nome:      "Paulo",
		sobrenome: "Silva",
	}

	fmt.Println(p1.getNomecompleto())

	p1.setNomeCompleto("Paulo Antônio")

	fmt.Println(p1)
}
