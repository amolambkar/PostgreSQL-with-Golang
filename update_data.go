package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func update() {
	psqlInfo := fmt.Sprintf("host = %s port= %d user = %s password = %s dbname = %s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sqlStatement := `
	UPDATE users SET "email"='amol@gmail.com',"age"=21 WHERE "id"=2`

	_, err = db.Exec(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}

	sqlStatement = `
	UPDATE users SET "first_name"=$1 WHERE "id"=$2`

	_, err = db.Exec(sqlStatement, "abhishek", 3)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(" record is updated")
	}

}
