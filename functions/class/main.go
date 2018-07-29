package main

import (
	"fmt"
)

var sum = func(a, b int) int {
	return a + b
}

var avg = func(num ...float64) float64 {
	all := 0.0
	for _, n := range num {
		all += n
	}

	return all / float64(len(num))
}

var listApproved = func(approved ...string) {
	fmt.Println("Lista de Aprovados")
	for i, appr := range approved {
		fmt.Printf("%d) %s\n", i, appr)
	}
}

func main() {
	fmt.Println(sum(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))

	approved := []string{"Mailonm", "Rafael", "Andrade", "Torres"}
	listApproved(approved...)

	fmt.Printf("Média %.2f", avg(1.2, 3.4, 1.7, 3.2, 6, 9.9, 9.9))

}
