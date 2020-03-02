package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBalisa()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar mais métodos
}

type bwm7 struct{}

func (b bwm7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bwm7) fazerBalisa() {
	fmt.Println("Baliza...")
}

func main() {
	var b esportivoLuxuoso = bwm7{}
	b.ligarTurbo()
	b.fazerBalisa()
}
