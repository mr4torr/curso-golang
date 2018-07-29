package main

import (
	"fmt"
	"time"
)

func isCousin(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func cousins(n int, c chan int) {
	ini := 2
	for i := 0; i < n; i++ {
		for cousin := ini; ; cousin++ {
			if isCousin(cousin) {
				c <- cousin
				ini = cousin + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c) // evita o deadlock
}

func main() {
	c := make(chan int, 30)

	go cousins(cap(c), c)

	for cousin := range c {
		fmt.Printf("%d ", cousin)
	}
	fmt.Println("Fim!")
}
