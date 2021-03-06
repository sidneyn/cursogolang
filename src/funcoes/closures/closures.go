package main

/**
Go suporta funções anônimas, que podem formar closures.
Funções anônimas são úteis quando você deseja definir uma função em linha sem ter que nomeá-lo.
Essa função intSeq retorna outra função, que definimos anônimamente no corpo do intSeq.
A função retornada se fecha sobre a variável i para formar o closure.

*/
import "fmt"

func intSeq() func() int {
	i := 0
	var funcao = func() int {
		i += 1
		return i
	}
	return funcao

}
func main() {
	nexInt := intSeq()

	fmt.Println(nexInt())
	fmt.Println(nexInt())
	fmt.Println(nexInt())
	fmt.Println(nexInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
