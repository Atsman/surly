package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS LINKS (
	hash char(7) NOT NULL,
	link TEXT NOT NULL,
	created_at timestamp DEFAULT NOW(),
	PRIMARY KEY(hash)
);
`

func InitDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=surly sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	return db
}
