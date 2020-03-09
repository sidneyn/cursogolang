package main

import "fmt"

func main() {
	/**** RANGE ****/
	// tres pontos indica que é um array e serve o compilador
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta por conta [...]! o numero correto de valores declarados 3pontos, sem ele você terá um slice!
	sum := 0
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
		sum += numero
	}
	fmt.Printf("SOMA %d ", sum)
	// ignorando o numero do indice _,
	for _, num := range numeros {
		fmt.Println(num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k, v := range kvs {
		fmt.Println("key", k, " valor", v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	/**
	referencia:
	https://gobyexample.com/range
	*/
}
