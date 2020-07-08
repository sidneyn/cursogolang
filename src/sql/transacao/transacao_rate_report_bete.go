package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	hostname      = "148.72.150.27"
	host_post     = 3254
	username      = "msbet"
	password      = "12345678"
	database_name = "bet_da_sorte"
)

func main() {

	pg_con_string := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, host_post, username, password, database_name)

	db, err := sql.Open("postgres", pg_con_string)

	if err != nil {
		panic(err)
		log.Fatal(err)
	}

	defer db.Close()

	// we can also ping our connection which will let us know if our connection is correct /// or not then we put an error-handling code right after that.
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("You are Successfully connected!")

	tx, err := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO rate_report(id, name) VALUES (?,?)")

	names := make([]string, 0)
	names = append(names, "Empate", "VencedorTime1", "VencedorTime2")

	for i := range names {
		_, err = stmt.Exec(i+500, names[i])
	}
	//	_, err = stmt.Exec(1, "Tiago")

	if err != nil {
		//	tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}
