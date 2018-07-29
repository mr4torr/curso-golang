package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // float64
	y := 2   // float32

	// convertendo valor flotuantes para um tipo
	fmt.Println("Float", (x / float64(y)))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println("Int", notaFinal)

	fmt.Println("Converter inteiro para string direto convert para o ASCII " + string(97))

	fmt.Println("Converter inteiro para string corretamente " + strconv.Itoa(123))

	num, _ := strconv.Atoi("123")

	fmt.Println("Alternativa para converter string para interiro ", num)
	fmt.Println("Teste de inteiro ", num-12)

	boolean, _ := strconv.ParseBool("fasd")
	if boolean {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}

}
