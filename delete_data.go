package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func delete() {
	psqlInfo := fmt.Sprintf("host = %s port= %d user = %s password = %s dbname = %s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sqlStatement := `
	DELETE FROM users WHERE "id"=2`

	_, err = db.Exec(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}

	sqlStatement = `
	DELETE FROM users WHERE "id"=$1`

	_, err = db.Exec(sqlStatement, 3)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(" record is deleted")
	}

}
