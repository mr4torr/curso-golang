package main

import "fmt"

func main() {
	var b byte = 3 // atribuição simples
	b++

	fmt.Println(b)
	i := 3 // inferencia de tipo
	i += 3 // atribuição aditiva
	i -= 3 // atribuição subtrativa
	i /= 2 // atribuição divisiva ( i = 1/2)
	i *= 2 // atribuição multiplicativa ( i = 1/2)
	i %= 2 // atribuição com mudulo ( i = 1 % 2)

	fmt.Println(i)

	x, y := 1, 2 // atribuição multiplo
	fmt.Println(x, y)

	x, y = y, x // invertendo valores
	fmt.Println(x, y)
}
