package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para o canal ( escrita )

	// <-ch // recebendo ddos do canal ( leitura )
	fmt.Println(<-ch)

	ch <- 2
	fmt.Println(<-ch)
}
