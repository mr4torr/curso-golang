package main

import (
	"fmt"
	"time"

	"gitlab.com/mr4torr/html"
)

func fastTitle(url1, url2, url3 string) string {
	t1 := html.GetTitle(url1)
	t2 := html.GetTitle(url2)
	t3 := html.GetTitle(url3)

	// select de controle é específica para concorrencia é
	// necessário tem um time em um dos case
	select {
	case t1 := <-t1:
		return t1
	case t2 := <-t2:
		return t2
	case t3 := <-t3:
		return t3
	// se diminuir o tempo ele retorna que todos perderam
	case <-time.After(time.Millisecond * 1000):
		return "Todos perderam!"
		// Se colocar um default ele não vai aguardar e vai sempre executar o default
		// default:
		// 	return "sem resposta ainda!"
	}
}

func main() {
	champion := fastTitle(
		"https://www.cod3r.com.br",
		"https://www.google.com",
		"https://www.amazon.com.br",
	)

	fmt.Println(champion)
}
