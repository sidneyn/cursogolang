package main

import (
	"fmt"
	m "math"
)

func main() {

	const PI float64 = 3.1415
	var raio = 3.2              //
	area := PI * m.Pow(raio, 2) // atribui e inicializa a variavel atraves :=
	fmt.Println("A área da circunferência é", area)
	fmt.Printf("A área da circunferência é %2f ", area)

	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa string!"
	fmt.Println(g, h, i)

}
