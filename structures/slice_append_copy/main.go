package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	var s []int

	s = append(s, 4, 5, 6)
	fmt.Println(s, a)

	s2 := make([]int, 2)
	copy(s2, s) // copy tem que ser slice
	fmt.Println(s2)

	s2[1] = 12

	fmt.Println(s2)

}
