package main

import "fmt"

type sport interface {
	turboOn()
}

type lux interface {
	automatic()
}

type sportAutomatic interface {
	sport
	lux
}

type bmw7 struct{}

func (b bmw7) turboOn() {
	fmt.Println("Turbo")
}

func (b bmw7) automatic() {
	fmt.Println("Automatico")
}

func main() {

	var b sportAutomatic = bmw7{}
	b.automatic()
}
