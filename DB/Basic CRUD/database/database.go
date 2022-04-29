package database

import (
	"database/sql"
)

func Connect() (*sql.DB, error) {
	connString := "root:root@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, error := sql.Open("mysql", connString)

	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		return nil, error
	}

	return db, nil

}
