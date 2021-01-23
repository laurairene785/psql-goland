package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//this is the function for the conection
func getConnection() *sql.DB {
	dsn := "postgres://goland:goland@192.168.0.1:5432/gocrud?sslmode=able"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db

}
