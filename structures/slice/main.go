package main

import (
	"fmt"
	"reflect"
)

func main() {

	a1 := [3]int{1, 2, 3} // Array
	s1 := []int{1, 2, 3}  // Slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[2:4]

	fmt.Println(s2)

	s3 := a2[:3]
	fmt.Println(s3)

}
