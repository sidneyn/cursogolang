package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:671160@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO usuarios(id, nome) VALUES (?,?)")

	nomes := make([]string, 0)
	nomes = append(nomes, "Bia", "Carlos", "Tiago")

	for i := range nomes {
		_, err = stmt.Exec(i+500, nomes[i])
	}
	//	_, err = stmt.Exec(1, "Tiago")

	if err != nil {
		//	tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
