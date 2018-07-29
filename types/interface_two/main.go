package main

import "fmt"

type sport interface {
	turboOn()
}

type ferrari struct {
	model    string
	turbo    bool
	velocity int
}

func (f *ferrari) turboOn() {
	f.turbo = true
}

func main() {
	// Exemplo sem chamada de interface
	car1 := ferrari{"F40", false, 0}
	car1.turboOn()
	car1.velocity = 100

	// Exemplo com chamada de interface
	var car2 sport = &ferrari{"F40", false, 0}
	car2.turboOn()

	fmt.Println(car1, car2)
}
