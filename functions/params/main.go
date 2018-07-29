package main

import "fmt"

func multiple(a, b int) int {
	return a * b
}

func exec(fnc func(int, int) int, p1, p2 int) int {
	return fnc(p1, p2)
}

func main() {
	result := exec(multiple, 3, 4)

	fmt.Println(result)
}
