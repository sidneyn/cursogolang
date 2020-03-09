package main

import "fmt"

func main() {
	// homogênea (mesmo tipo) e estática (fixo)
	var notas [3]float64
	notas[1] = 100
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	//notas[3] = 10
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	// calcular a media
	media := total / float64(len(notas))

	fmt.Printf("Total %.2f\n", total)
	fmt.Printf("Média %.2f\n", media)
	fmt.Println("total array", len(notas))

	tl := float64(len(notas))
	fmt.Println(tl)

	var total2 float64 = 0
	for _, value := range notas {
		total2 += value
	}
	fmt.Println(total2 / float64(len(notas)))

	x := [5]float64{
		//	98,
		93,
		//77,
		82,
		//83,
	}

	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
}
