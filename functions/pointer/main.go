package main

import (
	"fmt"
)

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// & É UTILIZADO PARA OBTER O ENDFEREÇO DA VARIAVEL
	inc2(&n) // por referencia
	fmt.Println(n)

}
