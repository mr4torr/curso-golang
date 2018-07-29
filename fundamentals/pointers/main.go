package main

import (
	"fmt"
)

func main() {
	i := 1

	var p *int = nil

	p = &i // pegando o endereço da variável

	*p++
	i++

	fmt.Println(p, *p, i, &i)

}
