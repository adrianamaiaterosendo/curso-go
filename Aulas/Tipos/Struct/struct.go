package main

import "fmt"

//colocar type + nome + struct
type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//função com receiver(receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	produto1 := produto{
		nome:     "banana",
		preco:    5,
		desconto: 0.1,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

}
