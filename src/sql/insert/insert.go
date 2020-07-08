package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type canal interface {
	make(chan byte)
}

func main() {
	db, err := sql.Open("mysql", "root:671160@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO usuarios(nome) VALUES(?)")
	if err != nil {
		panic(err)
	}
	nomes := make([]string, 0)

	nomes = append(nomes, "Maria", "João", "Sidney", "Marlene")

	for i := range nomes {

		// executando o sql e mostrando erro quando existir
		res, err := stmt.Exec(nomes[i])
		if err != nil {
			panic(err)
		}

		fmt.Sprintf("inserido usuario", nomes[i])

		// ultimo inserido
		id, _ := res.LastInsertId()
		fmt.Println("ultimo id inserido no banco ", id)
		fmt.Println()

	}

	/*	stmt.Exec("João")
		/*
			res, _ := stmt.Exec("Pedro")

			id, _ := res.LastInsertId()
			fmt.Println(id)

			linhas, _ := res.RowsAffected()
			fmt.Println(linhas)
	*/
}
