package main

import (
	"fmt"

	"gitlab.com/mr4torr/html"
)

func send(origin <-chan string, destination chan string) {
	for {
		destination <- <-origin
	}
}

// multiplexar = misturar mensages de canais em um unico
func group(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go send(input1, c)
	go send(input2, c)
	return c
}

func main() {
	c := group(
		html.GetTitle("https://www.cod3r.com.br", "https://www.google.com"),
		html.GetTitle("https://www.amazon.com.br", "https://www.youtube.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
