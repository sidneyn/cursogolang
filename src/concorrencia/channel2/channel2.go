package main

// Channel (canal) - -e a forma de comunicação entre goroutines ponto de sincronização ele fica esperando o dado chegar
// channel é um tipo  como qualquer outro tipo na linguagem, como inteiro, string boolean
import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal
	time.Sleep(time.Second)
	c <- 3 * base
	time.Sleep(time.Second)
	c <- 4 * base
	time.Sleep(time.Second)
	c <- 5 * base
}

func main() {

	c := make(chan int)

	go doisTresQuatroVezes(2, c)
	fmt.Println("A")
	//e, f, h := <-a, <-b, <-d // < recebendo dados do canal
	fmt.Println(<-c, <-c, <-c)

	fmt.Println("B")
	//fmt.Println(<-c)

}
