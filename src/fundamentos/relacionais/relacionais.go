package main

import (
	"fmt"
	"time"
)

func main() {
	/*** OPERADORES RELACIONAIS ****/
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 1600)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome      string
		Sobrenome string
	}

	type PessoaFisica struct {
		cpf string
	}

	p1 := Pessoa{"João", "bosco"}
	pf1 := PessoaFisica{"02135144436"}

	p2 := Pessoa{"João", "boco"}
	fmt.Println("Pessoas:", p1 == p2)
	fmt.Println("Pessoas Fisica:", pf1)
	fmt.Println("Pessoas Fisica:", p1)

}
