package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1> Hora certa: %s<h1>", s)
	fmt.Fprintf(w, "<div>DIV 1</div>")

}

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal
	time.Sleep(time.Second)
	c <- 3 * base
	time.Sleep(time.Second)
	c <- 4 * base
	time.Sleep(time.Second)
	c <- 5 * base
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Execuntando...")
	log.Fatalln(http.ListenAndServe(":3000", nil))
}
