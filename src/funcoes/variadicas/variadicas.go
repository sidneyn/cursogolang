package main

import "fmt"

/**
Funções variádicas podem ser chamadas com qualquer número de argumentos à direita. Por exemplo, fmt.Println é uma função variádica comum.
Aqui temos uma função que irá pegar um número arbitrário de inteiros como argumentos.
*/
func soma(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func somaNotas(notas ...int) {
	fmt.Println(notas, " ")
	qtd := len(notas)
	soma_notas := 0
	for _, nota := range notas {
		soma_notas += nota

	}
	media := soma_notas / qtd
	fmt.Println("Media ", media)
}

// soma com float variadicas
func somaFloat(valores ...float64) float64 {

	var total float64 = 0
	for _, valor := range valores {
		total += valor

	}
	qtd := len(valores)
	res := total / float64(qtd)
	return res

	//fmt.Printf("total %.2f ", total / float64(qtd))
	// fmt.Println("total ", total / float64(qtd))
}
func resultadoNotas(res float64) string {

	switch res {
	case 10.0:
		return "Parabéns aprovado com nota maxima!"
	case 9.9, 8.0:
		return "Parabéns aprovado acima media!"
	case 7.9, 6.0:
		return "Parabéns aprovado acima media!"
	default:
		return "sem valor"
	}
}
func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	soma(1, 2)
	soma(1, 2, 3)

	nums := []int{1, 2, 3, 4}

	soma(nums...)

	somaNotas(6, 6, 7)

	rs := somaFloat(9.5, 8.6, 9.9)
	fmt.Println("somatorio", rs)
	rest := resultadoNotas(rs)
	fmt.Println(rest)

	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	imprimirAprovados(aprovados...)

}
