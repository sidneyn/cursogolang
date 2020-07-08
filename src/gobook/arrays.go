package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 15
	x[1] = 24
	x[2] = 31
	x[3] = 33
	x[4] = 100
	//fmt.Println(x)
	//fmt.Println(len(x))
	y := []float64{98, 93, 77}//82,
	//83

	fmt.Println(y)

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	var total float64 = 0
	for _, value := range x {
		total += float64(value)
	}
	fmt.Println(total / float64(len(x)))

}
