package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//Slice não é um array! Slide define um pedaço de um array.
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2] // novo slice, mas aponta para o mesmo array
	fmt.Println(a2, s3)

	/********************************/

	/*
			Uma fatia é um segmento de uma matriz. As matrizes, como fatias, são indexáveis ​​e têm um comprimento.
			Diferentemente das matrizes, esse tamanho pode mudar. Aqui está um exemplo de uma fatia:

			var x []float64

			Se você deseja criar uma fatia, deve usar a makefunção interna:
			x := make([]float64, 5)


		 	Isso cria uma fatia associada a uma float64 matriz subjacente de comprimento 5.
		  	As fatias estão sempre associadas a alguma matriz e, embora nunca possam ser
			maiores que a matriz, podem ser menores. A makefunção também permite um terceiro parâmetro:

			x: = make ([] float64, 5, 10)

			Outra maneira de criar fatias é usar a [low : high]expressão:

			arr: = [5] float64 {1,2,3,4,5}
			x: = arr[0:5]


			low é o índice de onde iniciar a fatia e highé o índice de onde finalizá-la (mas não incluindo o próprio índice).
			Por exemplo, enquanto arr[0:5]retorna [1,2,3,4,5], arr[1:4]retorna [2,3,4].
			Por conveniência, também podemos omitir low, highou mesmo ambos lowe high. arr[0:]é o mesmo que arr[0:len(arr)],
			arr[:5]é o mesmo que arr[0:5]e arr[:]é o mesmo que arr[0:len(arr)].
	*/

}
