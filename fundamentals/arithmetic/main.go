package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10
	b := 8
	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplacação =>", a*b)
	fmt.Println("Operador Módulo =>", a%b) // tras o resto da divição 3/2 = 1 4/2 = 0

	fmt.Println("&& =>", a&b)  // bitwise ( calculo bit a bit 3(11)& 2(10) = 10)
	fmt.Println("OR =>", a|b)  // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01

	fmt.Println("Maior => ", math.Max(float64(a), float64(b))) // pega o valor maximo

	c := 4.0
	d := 3.0

	fmt.Println("Exponenciação => ", math.Pow(c, d)) // = (4*4)*4
	fmt.Println("Menor => ", math.Min(c, d))         // pega o valor minimo

}
