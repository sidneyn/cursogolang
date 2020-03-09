package main

import (
	"fmt"
)

func main() {
	p1, p2 := Ponto{2.8, 2.0}, Ponto{2.8, 4.0}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))

}
