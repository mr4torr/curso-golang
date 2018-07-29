package main

import (
	"fmt"
)

type item struct {
	productID int
	qtd       int
	price     float64
}

type order struct {
	id     int
	userID int
	itens  []item
}

// Método com função de receiver ( receptor )
func (o order) priceTotal() float64 {
	all := 0.0
	for _, item := range o.itens {
		all += item.price * float64(item.qtd)
	}
	return all
}

func main() {
	// var product1 product

	order := order{
		userID: 1,
		itens: []item{
			item{productID: 4, qtd: 5, price: 23.45},
			item{1, 2, 12.10},
			item{12, 1, 23.10},
			item{11, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do pedido é %.2f", order.priceTotal())

}
