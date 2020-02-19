package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	/**** OPERADORES LOGICOS ****/
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("TV50: %t, Tv32: %v, Sorvete: %t, Saudável: %t \n", tv50, tv32, sorvete, !sorvete)
}
