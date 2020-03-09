package main

import (
	"math"
)

// iniciando com letra maiuscula é PUBLICA (visivel e fora do pacote)!
// iniciando com letra minuscula é PRIVADO (visivel dentro  do pacote)!

// Por exemplo
// Ponto -gerará algo público
// ponto ou _Ponto - gerará algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
