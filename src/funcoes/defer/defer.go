package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!")
	defer fmt.Println("fim2!")
	if numero > 500 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	//fmt.Println(obterValorAprovado(6000))
	//fmt.Println(obterValorAprovado(3000))
	fmt.Println(obterValorAprovado(30))

}
