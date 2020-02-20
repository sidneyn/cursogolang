package main

import "fmt"

/**
Mapas
Um mapa é uma coleção não ordenada de pares de valores-chave.
Também conhecidos como uma matriz associativa, uma tabela de hash ou um dicionário, os mapas são usados ​​para procurar um valor por sua chave associada.
Aqui está um exemplo de mapa no Go:

var x map [string] int

O tipo de mapa é representado pela palavra-chave map, seguida pelo tipo de chave entre colchetes e, finalmente, pelo tipo de valor.
Se você lesse isso em voz alta, diria " xé um mapa de strings para ints".
Mapas de matrizes e fatias podem ser acessados ​​usando colchetes. Tente executar o seguinte programa:
*/
func main() {
	//var x map [string] int
	//x["chave"] = 10
	//fmt.Println(x)
	x := make(map[string]int)
	x["chave"] = 10
	fmt.Println(x["chave"])

	// Vejamos um exemplo de programa que usa um mapa:
	elementos := make(map[string]string)
	elementos["H"] = "Hidrogênio"
	elementos["He"] = "Hélio"
	elementos["Li"] = "litio"
	elementos["Be"] = "Berilio"
	elementos["B"] = "Boro"
	elementos["C"] = "Carbono"
	elementos["N"] = "Nitrogênio"
	elementos["O"] = "Oxigênio"
	elementos["F"] = "Flúor"
	elementos["Ne"] = "Neon"

	fmt.Println(elementos["Li"])
	fmt.Println(elementos["Un"])

	nome, ok := elementos["H"]
	nome2, ok := elementos["Un"]

	fmt.Println(nome, ok)
	fmt.Println(nome2, ok)

	// Criar mapas curtos
	/*
		elementos2 := map[string] string {
			"H" : "Hidrogênio",
			"He" : "Hélio",
			"Li" : "litio",
			"Be" : "Berilio",
			"B" : "Boro",
			"C" : "Carbono",
			"N" : "Nitrogênio",
			"O" : "Oxigênio",
			"F" : "Flúor",
			"Ne" : "Neon",
		}
		fmt.Println(elementos)*/
}
