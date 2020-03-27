package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	fmt.Println("falando")
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale ("Maria", "Pq vc não fala comigo", 3)
	// fale ("João", "Só posso falar depois de voce!" ,1)

	go fale("Maria", "Ei...", 500)
	go fale("João", "Opa...", 500)
	fmt.Printf("FIM ")
	//go fale ("Maria", "Entendi!!!", 10)
	// fale ("Maria", "Entendi!!!", 20)
}
