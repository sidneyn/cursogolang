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

	// UPDATE TRANSACAO
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")

	_, err = stmt.Exec("Uóxiton Istive", 1)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// DELETE TRANSACAO
	stmt2, _ := tx.Prepare("DELETE FROM usuarios WHERE id = ?")
	_, err = stmt2.Exec(3)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
}
