package main

import "fmt"

func main() {

	fmt.Println("PRINTF")
	var i interface{} = 23
	fmt.Printf("%v\n", i)

	fmt.Println("PRINT")
	fmt.Print("mesma ")
	fmt.Print(" linha ")
	fmt.Println("PRINTFLN")
	fmt.Println("    Primeira linha de texto")
	fmt.Println("    Segunda linha de texto")
	//fmt.Fprint("Texto no Fprint %f", "texto")

}
