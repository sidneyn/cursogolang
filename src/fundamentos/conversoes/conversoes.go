package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Teste" + string(97))

	// convertendo int para string
	fmt.Println("Teste converte " + strconv.Itoa(123))

	// convertendo string para int
	num, err := strconv.Atoi("123")
	fmt.Println(num - 122)
	fmt.Println(err)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("false")
	}

	// convertendo Strings and Bytes
	e := "my string é tempo"

	f := []byte(e)

	g := string(f)

	fmt.Println(e)

	fmt.Println(f)

	fmt.Println(g)

}
