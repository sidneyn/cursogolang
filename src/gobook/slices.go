package main

import "fmt"

func main() {
	//var x []float64

	x := make([]float64, 5)
	y := make([]float64, 5, 10)

	fmt.Println(x)
	fmt.Println(y)
}
