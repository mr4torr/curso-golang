package main

import "fmt"

func closure() func() {
	x := 10
	var fnc = func() {
		fmt.Println(x)
	}
	return fnc
}

func main() {
	x := 20
	fmt.Println(x)

	printX := closure()
	printX()
}
