package main

import (
	"fmt"
)

type print interface {
	toString() string
}

type person struct {
	fisrtName string
	lastName  string
}

type product struct {
	name  string
	price float64
}

func (p person) toString() string {
	return p.fisrtName + " " + p.lastName
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func printer(x print) {
	fmt.Println(x.toString())
}

func main() {
	var nameVar print = person{"Mailon", "Torres"}
	fmt.Println(nameVar.toString())
	printer(nameVar)

	nameVar = product{"Calça Jean", 79.90}
	fmt.Println(nameVar.toString())
	printer(nameVar)

}
