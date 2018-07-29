package main

import "fmt"

func getAmountApproved(num int) int {
	// Ele executa um pouco antes de fazer o return
	defer fmt.Println("fim!")

	if num >= 5000 {
		fmt.Println("Valor alto.... ")
		return 5000
	}

	fmt.Println("Valor baixo.... ")
	return num
}

func main() {
	fmt.Println(getAmountApproved(5000))
}
