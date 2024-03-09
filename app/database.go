package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

func NewDB() *sql.DB {
	dbInfo := "host=localhost port=5432 user=postgres password=ArvaLapak1209 dbname=golang_todo_list sslmode=disable"
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	//db pooling
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
