package main

import (
	"fmt"
	"strings"
)

type person struct {
	firstName string
	lastName  string
}

// Método com função de receiver ( receptor )
func (p *person) getFullName() string {
	return p.firstName + " " + p.lastName
}

func (p *person) setFullName(fullName string) {
	parts := strings.Split(fullName, " ")
	p.firstName = parts[0]
	p.lastName = parts[1]
}

func main() {
	// var product1 product

	person1 := person{"Mailon", "Torres"}
	fmt.Println(person1.getFullName())
	fmt.Println(person1)

	person1.setFullName("Zuliana Lima")
	fmt.Println(person1.getFullName())
	fmt.Println(person1)

}
