package main

import "fmt"

func change(p1, p2 int) (seg int, pri int) {
	seg = p1
	pri = p2
	return // chamado de retorno limpo, traz o que foi definido na definição da função
}

func main() {
	x, y := change(2, 3)
	fmt.Println(x, y)

	seg, pri := change(7, 1)
	fmt.Println(seg, pri)
}
