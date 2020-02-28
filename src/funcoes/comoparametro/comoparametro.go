package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func adicao(c, d int) int {
	return c + d
}

func subtracao(e, f int) int {
	return e - f
}

func main() {
	resultado := exec(multiplicacao, 3, 4)
	soma := exec(adicao, 10, 12)
	subtracao := exec(subtracao, 24, 11)
	fmt.Println(resultado)
	fmt.Println(soma)
	fmt.Println(subtracao)
}
