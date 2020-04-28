package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
	Categ []categoria
}
type categoria struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

/*
func getJson (p produto) <-chan string {

	resJson := make(chan string)
	p1 := produto{1, "Notebook", 189.90, []string{"Promoção", "Eletronico"}, []categoria{{ID: 10, Nome: "fralda"}, {ID: 11, Nome: "perfumes"}}}

	return resJson
}
*/

func main() {
	//struct para json
	type can1 interface{}
	var c can1 = produto{1, "Notebook", 189.90, []string{"Promoção", "Eletronico"}, []categoria{{ID: 10, Nome: "fralda"}, {ID: 11, Nome: "perfumes"}}}
	fmt.Println(c)

	//can1 :=

	p1Json, _ := json.Marshal(p1) // metodo que retorna dois valores
	fmt.Println(p1Json)
	fmt.Println(string(p1Json))

	// json para struct

	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":8.90, "tags": ["Papelaria", "Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2.Nome)
	fmt.Println(p2.Preco)
	fmt.Println(p2.Tags[0])
	fmt.Println(p2.Tags[1])

	var c2 categoria
	jsonString2 := `{"id":19,"nome":"Papelaria Nova"}`
	json.Unmarshal([]byte(jsonString2), &c2)

	fmt.Println(c2.ID)
	fmt.Println(c2.Nome)
}
