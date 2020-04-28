package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:671160@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "CREATE database IF NOT EXISTS cursogo")
	exec(db, "USE cursogo")
	exec(db, "DROP TABLE IF EXISTS usuarios")
	exec(db, `CREATE TABLE usuarios (
					id INTEGER AUTO_INCREMENT,
					nome VARCHAR (80),
					PRIMARY KEY(id)
	)`)
}
