package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// funcao tras um ponteiro do tipo int
func zeroptr(iptr *int) {
	*iptr = 10
}

func main() {
	i := 1
	fmt.Println("inicial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// A sintaxe &i fornece o endereço da memória do i, ou seja, um ponteiro para i.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("ponteiro:", &i)
}

/**
referencia:
http://goporexemplo.golangbr.org/pointers.html
*/
