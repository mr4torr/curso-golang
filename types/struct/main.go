package main

import (
	"fmt"
)

type product struct {
	name     string
	price    float64
	discount float64
}

// Método com função de receiver ( receptor )
func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	// var product1 product

	var product1 = product{
		name:     "Lapis",
		price:    1.79,
		discount: 0.05,
	}

	fmt.Println(product1, product1.priceWithDiscount())

	var product2 = product{"Caderno", 9.79, 0.05}
	fmt.Println(product2, product2.priceWithDiscount())

}
