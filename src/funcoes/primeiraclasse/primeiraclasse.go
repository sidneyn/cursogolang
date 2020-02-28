package main

import "fmt"

var soma = func(a, b, c int) (int, int) {
	return a + b, a + b - c
}

func main() {
	fmt.Println(soma(2, 3, 2))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
