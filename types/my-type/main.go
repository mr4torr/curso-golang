package main

import (
	"fmt"
)

// Cria um novo typo com receiver
type note float64

func (x note) between(ini, fin float64) bool {
	return float64(x) >= ini && float64(x) <= fin
}

func noteWith(n note) string {
	switch {
	case n.between(9.0, 10.0):
		return "A"
	case n.between(7, 8.99):
		return "B"
	case n.between(5, 6.99):
		return "C"
	case n.between(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(noteWith(9.8))
	fmt.Println(noteWith(6.9))
	fmt.Println(noteWith(2.1))
}
