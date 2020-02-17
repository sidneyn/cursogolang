package main

import "fmt"

// Não tem operador ternario
func obterResultado(nota float64) string {
	// return nota >= 6  "aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovador"
}

// executando
func main() {
	fmt.Println(obterResultado(6.2))
}

/**
executando build
go build file.go
executando dev checar erro no codigo
go dev file.go
executa run compila e executa sem criar build
go run file.go
limpa file build criado
go clean file.go

*/
