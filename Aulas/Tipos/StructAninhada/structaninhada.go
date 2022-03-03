package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	item   []item
}

func (p pedido) totalPedido() float64 {
	total := 0.0
	for _, item := range p.item {
		total += (float64(item.qtde) * item.preco)
	}
	return float64(total)
}
func main() {

	pedido := pedido{
		userID: 1,
		item: []item{
			{
				produtoID: 1,
				qtde:      10,
				preco:     5.50,
			},
			{
				produtoID: 2,
				qtde:      1,
				preco:     1,
			},
		},
	}

	fmt.Println(pedido.totalPedido())
}
