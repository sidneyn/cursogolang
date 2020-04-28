package main

import "fmt"

func main() {
	var i int
	array1 := []int{1, 2, 3}
	var slice1 []int
	var nomes []string
	array1 = append(array1, 4, 5, 6)
	slice1 = append(slice1, 4, 5, 6)
	nomes = append(nomes, "Sidney", "Tiago", "Thomaz")

	var loteria []string
	loteria = append(loteria, "caixa", "pernambuco", "ceara")

	fmt.Println(array1, slice1, nomes, "total de nomes ", len(nomes))

	for i = 0; i < len(nomes); i++ {
		fmt.Println("Nomes:  ", nomes[i], " posicao ", i, " / ", len(nomes))

	}
	for _, nome := range nomes {
		fmt.Println("NOME: ", nome)

	}

	for _, loteria := range loteria {
		fmt.Println("loteria ", loteria)
	}

}
