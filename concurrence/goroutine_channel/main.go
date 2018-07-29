package main

import (
	"fmt"
	"time"
)

// Channel - é a forma de comunicação entre goroutine
// Channe é um tipo

func calcParalela(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base

	time.Sleep(3 * time.Second)
	c <- 5 * base
}

func main() {
	c := make(chan int)
	go calcParalela(2, c)
	fmt.Println("A")

	a, b, d := <-c, <-c, <-c // recebendo dados do canal
	fmt.Println("B")

	fmt.Println(a, b, d)

	fmt.Println(<-c)

	// Se adicionar mais retorno de channel ele reporta deadlock
	// fmt.Println(<-c)

}
