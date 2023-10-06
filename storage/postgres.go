package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	connstr := fmt.Sprintf("user:%s password:%s dbname:%s sslmode:disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		panic(err)
	}

	DB = db

}

func CreateTables() {
	_, err := DB.Exec(
		`create table if not exists apod_metadata (
			date date primary key,
			title text,
			explanation text,
		);
		`)

	if err != nil {
		panic(err)
	}
}
