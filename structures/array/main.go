package main

import "fmt"

func main() {
	numeros := [...]float64{1, 2, 3, 4, 5}

	// Com o index
	for i, numero := range numeros {
		fmt.Printf("%d) %.2f\n", i+1, numero)
	}

	// Ignorar o index
	for _, numero := range numeros {
		fmt.Printf("%.2f\n", numero)
	}
}
