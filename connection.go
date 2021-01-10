package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//this is the function for the conection
func getConnection() *sql.DB {
	dsn := "postgres://go:go@127.0.0.1:5432/goweb?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db

}
