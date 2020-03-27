package main

import (
	"cursogolanghtml/html"
	"fmt"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - mistura (mensagens num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://wwww.google.com", "https://www.google.com"),
		html.Titulo("https://wwww.amazon.com", "https://www.amazon.com"),
	//html.Titulo("https://wwww.cod3r.com.br", "https://www.google.com"),
	//	html.Titulo("https://wwww.amazon.com", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	//fmt.Println(<-c, "|", <-c)
}
