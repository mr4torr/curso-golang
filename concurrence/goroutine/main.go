package main

import (
	"fmt"
	"time"
)

func speak(person, text string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s:%s (interação %d)\n", person, text, i+1)
	}
}

func main() {
	// Exemplo 1
	// Modelo serial
	// speak("Maria", "Pq vc não fala comigo?", 3)
	// speak("João", "Só posso falar depois de vc!", 1)

	// Exemplo 2
	// Modelo paralelo
	// go speak("Maria", "Ei...", 5)
	// go speak("João", "opa...", 5)
	/*
		Como a função demos para ser imprimida, faz com que não conclua sua
		execução, já que a traide principal terminou antes de executar
	*/
	// Colocado essa função para executar as goruntine
	// time.Sleep(time.Second * 5)

	// Exemplo 3
	go speak("Maria", "Entendi!!!", 10)
	speak("João", "Parabéns", 5)

}
