package main

import (
	"fmt"
	"time"
)

func rotine(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("Só depois que foi lido")
}
func main() {
	c := make(chan int) // canal sem buffer

	go rotine(c)

	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock = não tem channel para ser retornado
	fmt.Println("Fim...")

}
