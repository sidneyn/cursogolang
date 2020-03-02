package main

import "fmt"

type curso struct {
	nome string
}

type materia struct {
	curso string
}

func main() {
	// interface do tipo generica vazia
	var qualquercoisa interface{}
	fmt.Println(qualquercoisa)

	qualquercoisa = 3
	fmt.Println(qualquercoisa)
	qualquercoisa = "Opa"
	fmt.Println(qualquercoisa)
	qualquercoisa = true
	fmt.Println(qualquercoisa)
	qualquercoisa = materia{"GoLang"}
	fmt.Println(qualquercoisa)

	var coisa interface{}
	fmt.Println(coisa)

	// passando um valor inteiro para interface
	coisa = 3
	fmt.Println(coisa)
	// criando outra interface generica nula
	type dinamico interface{}

	// passando um valor string para a interface dinamica
	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	// passando um boolean para a interface coisa2
	coisa2 = true
	fmt.Println(coisa2)

	// atribuindo um struct a variavel coisa2
	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2)

}
