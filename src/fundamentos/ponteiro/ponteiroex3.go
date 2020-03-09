package main

import "fmt"

func zero(x int) {
	x = 0
}
func zero2(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5
	/*
		Neste programa, a função zero não modificará o
		variável x original na função principal. Mas e se nós
		queria? Uma maneira de fazer isso é usar dados especiais
		tipo conhecido como ponteiro:
	*/
	y := 5
	zero2(&y)
	fmt.Println(y) // x is 0
}
