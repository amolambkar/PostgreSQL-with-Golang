package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func insert() {
	psqlInfo := fmt.Sprintf("host = %s port= %d user = %s password = %s dbname = %s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sqlStatement := `
	INSERT INTO users ("age", "first_name", "last_name","email")
	VALUES (20, 'Amol', 'Ambkar','amolambkar00@gmail.com')`

	_, err = db.Exec(sqlStatement)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New record is added")
	}

}
