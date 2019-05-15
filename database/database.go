package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Client *sqlx.DB

func InitDB() {
	db, err := sqlx.Connect("postgres", "user=ubilling password=pbilling dbname=billing host=db port=5432 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	Client = db
}
