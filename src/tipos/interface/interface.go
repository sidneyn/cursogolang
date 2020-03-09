package main

import (
	"errors"
	"fmt"
)

type imprimivel interface {
	toString() string
}

type corrigivel interface {
	validaNome(nome string) (string, error)
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}
type pessoa1 struct {
	nome string
}

func (p pessoa1) validaNome(nome string) (string, error) {

	if nome == "" {
		return string("-1"), errors.New("campo inválido!")
	}
	return p.nome, nil
}

/**
 * Conceito de Interface são implementados implicitamente no GO de forma que ao criar um tipo interface e dentro dele incluir dentro dele um
 * metodo (func), atributo (attributes) ou construtor (struct) a funcionalidade poderá ser usada conforme a necessidade.
 *
 */

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}
func corrigir(y corrigivel) {
	fmt.Println(y.validaNome("sidney"))
}

func main() {
	// a inferface utiliza o conceito de polimofismo Orientação a Objeto
	// a variavel coisa é do tipo pessoa
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Calça Jeans", 179.90}
	imprimir(p2)

	// Corrigir pedido do cliente

	//	var nome  corrigivel = pessoa1{"Sidney "}
	var nome2 corrigivel = pessoa1{""}
	var nomesobrenome corrigivel = pessoa1{""}
	//fmt.Println(nome)
	fmt.Println(nome2)
	fmt.Println(nomesobrenome)
	//corrigir(nome)

}
