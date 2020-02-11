package main

import "fmt"

func main() {

	fmt.Println("PRINTF")
	var i interface{} = 23
	fmt.Printf("%v\n", i)

	fmt.Println("PRINT")
	fmt.Print("mesma ")
	fmt.Print(" linha ")
	fmt.Println("    Primeira linha de texto")
	fmt.Println("    Segunda linha de texto")

	x := 3.151516

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	fmt.Printf("O valor de x é %.2f :::", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, c, d)
	fmt.Printf("\n%v %v %v %s ", a, b, c, d)

}
