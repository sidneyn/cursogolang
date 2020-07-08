package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("numero par ", i)
		} else {
			fmt.Println("numero impar ", i)
		}

		if i == 0 {
			fmt.Println("Zero")
		} else if i == 1 {
			fmt.Println("One")
		} else if i == 2 {
			fmt.Println("Two")
		} else if i == 3 {
			fmt.Println("Three")
		} else if i == 4 {
			fmt.Println("Four")
		} else if i == 5 {
			fmt.Println("Five")
		}
	}
}
