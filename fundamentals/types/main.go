package main

import (
	"fmt"
	"reflect"
)

func main() {
	// String
	i2 := 123
	fmt.Println("O int é", reflect.TypeOf(i2))

	n1 := float32(49)
	fmt.Println("O literal é", reflect.TypeOf(n1))

	n2 := float64(49.90)
	fmt.Println("O literale é", reflect.TypeOf(n2))

	n3 := 49.90
	fmt.Println("O literal é", reflect.TypeOf(n3))

	n4 := 49
	fmt.Println("O literal é", reflect.TypeOf(n4))

	b1 := true
	fmt.Println("O tipo é boolean", reflect.TypeOf(b1))

	s1 := "O tipo string"
	fmt.Println(s1 + "adicionado string")

	s2 := " O tipo string "
	fmt.Println(s1, len(s2))

	s3 := `String
	com
		crase`
	fmt.Println(s3, len(s3))

	c1 := 'a' // char é com aspas simples e somente um catracter retorna um int32
	fmt.Println("Tipo char = ", reflect.TypeOf(c1))
}
