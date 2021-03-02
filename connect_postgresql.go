package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "amol"
	password = "Pass@123"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host = %s port= %d user = %s password = %s dbname = %s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Pong")
	}

	insert()
}
