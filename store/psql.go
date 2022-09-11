package store

import (
	"database/sql"
	"log"
	"mini-wallet/libs"

	_ "github.com/lib/pq"
)

func InitPsql(config libs.Config) *sql.DB {
	db, err := sql.Open("postgres", config.DBSource)
	if err != nil {
		log.Fatalf("error, cannot connect to DB: %s", err.Error())
	}

	return db
}
