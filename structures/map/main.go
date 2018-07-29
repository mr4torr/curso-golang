package main

import (
	"fmt"
)

func mapSimple() {
	itens := make(map[int]string)

	itens[123] = "Maria"
	itens[321] = "Joao"
	itens[231] = "Jose"

	fmt.Println(itens)
	delete(itens, 321)
	delete(itens, 32131)

	for key, value := range itens {
		fmt.Printf("%d => %s\n", key, value)
	}
}

func mapDeclado() {
	// fecha o map e não permite adicionar mais nada
	funcAmount := map[string]float64{
		"João José":      11325.45,
		"Grabrile Silva": 15456.78,
		"Pedro":          1500.28}

	// con firgula no final permite adicionar mias map
	funcAmount2 := map[string]float64{
		"João José":      11325.45,
		"Grabrile Silva": 15456.78,
		"Pedro":          1500.28,
	}

	funcAmount2["Victor"] = 1324.45

	fmt.Println(funcAmount, funcAmount2)

}

func mapAlinhado() {
	funcAmount := map[string]map[string]float64{
		"J": {
			"João José": 11325.45,
		},
		"G": {
			"Grabrile Silva": 15456.78,
		},
		"P": {
			"Pedro": 1500.28,
		},
	}

	fmt.Println(funcAmount)

	for initial, funcs := range funcAmount {

		fmt.Println(initial, funcs)
	}
}

func main() {
	// mapSimple()
	// mapDeclado()
	mapAlinhado()
}
