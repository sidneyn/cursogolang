package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second) // tempo da função dormir
	c <- 1                  // operacao bloqueante
	//fmt.Println("Só depois que foi lido")
}

func rotina2(s chan string) {

	time.Sleep(time.Second) // tempo para ver a função rodando
	s <- "Sidney"           // atribuindo dado ao canal
	s <- "Nogueira"         // atribuindo dado ao canal
}

func main() {

	c := make(chan int) // canal 1 sem buffer

	s := make(chan string) // canal 2

	go rotina(c)
	// enviando para o canal
	fmt.Println(<-c) // recebe o dado do canal, operacao bloqueante
	fmt.Println("Foi lido")
	// abrindo um nova thred com goroutine
	go rotina2(s)

	nome, sobrenome := <-s, <-s
	fmt.Println("bom dia ", nome, sobrenome)

	//fmt.Println(<-s) // deadlock  fatal error: all goroutines are asleep - deadlock!
	fmt.Println("Fim")

}
