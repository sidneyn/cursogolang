package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Println("calculando...")

	output := input * 2

	fmt.Println(output)
	//log.Println("uso do scanf ")
}
