package main

import "fmt"

// Não existe operador ternario no GO lang
// true == false ? 'sim' : "não"

func returnResult(nota float64) string {

	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(returnResult(6.2))
}
