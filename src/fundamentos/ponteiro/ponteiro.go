package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	p = &i // pegando o endereço da variavel
	*p++   // desreferenciando (pegando o valor)
	i++

	// GO nao tem aritmetica de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)

}
