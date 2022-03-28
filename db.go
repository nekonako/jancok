package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func newDb() *sql.DB {
	connStr := "postgres://postgres:@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxIdleTime(10)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(-1)
	return db
}
