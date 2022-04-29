package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	urlConn := "root:root@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", urlConn)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	defer db.Close() // Closing the connection as late as possible

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database is connected and available!")
	}

	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close() // Closing the rows cursor as late as possible

	fmt.Println(err)

}
