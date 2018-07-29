package main

import "fmt"

func fat(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatAnt, _ := fat(n - 1)
		return n * fatAnt, nil
	}
}

func fatCompact(n uint) uint {
	if n == 0 {
		return 1
	}

	return n * fatCompact(n-1)
}

func main() {
	result, _ := fat(5)
	fmt.Println(result)

	_, err := fat(-5)
	if err != nil {
		fmt.Println(err)
	}

	result2 := fatCompact(5)
	fmt.Println(result2)

}
