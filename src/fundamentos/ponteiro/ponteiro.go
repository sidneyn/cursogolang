package main

import "fmt"

func main() {
	i := 1

	// * ponteiro é uma referencia de memoria
	// nil é equivalete a null
	var p *int = nil
	// pege o endereço de memoria e atribuir ao ponteiro
	p = &i // pegando o endereço da variavel
	*p++   // desreferenciando (pegando o valor)
	i++

	// GO nao tem aritmetica de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)

}
