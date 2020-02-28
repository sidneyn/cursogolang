package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}
type produtoSemDesconto struct {
	nome  string
	preco float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}
func (p2 produtoSemDesconto) precoSemDesconto() produtoSemDesconto {
	return p2
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	var produtoSD produtoSemDesconto
	produtoSD = produtoSemDesconto{
		nome:  "notebook2",
		preco: 1.59009,
	}
	fmt.Println(produto1, produto1.precoComDesconto())
	fmt.Println(produtoSD, produtoSD.precoSemDesconto().nome, produtoSD.precoSemDesconto().preco)

	produto2 := produto{"Notebook", 1989.90, 0.10}

	fmt.Println(produto2.precoComDesconto())
}
