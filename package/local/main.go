package main

import "fmt"

func main() {
	p1 := Ponto{2, 2}
	p2 := Ponto{2, 9}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
