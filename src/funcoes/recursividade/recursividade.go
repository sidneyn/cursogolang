package main

import (
	"fmt"
)

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		fmt.Printf("fatorial %v n %v ", fatorialAnterior, n)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(2)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
