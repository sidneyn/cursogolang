package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// bitwise
	fmt.Println("E =>", a&b)
	fmt.Println("Ou =>", a|b)
	fmt.Println("Xou =>", a^b)

	c := 3.0
	d := 2.0

	// outra operacoes usando match...
	fmt.Println("Maior =>", math.Max(float64(c), float64(d)))
	fmt.Println("Menor =>", math.Min(float64(c), float64(d)))
	fmt.Println("Exponenciação =>", math.Pow(c, d))

}
