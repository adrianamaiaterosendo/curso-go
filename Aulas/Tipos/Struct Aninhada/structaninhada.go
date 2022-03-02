package main

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
