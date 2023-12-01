package database

import (
	"csv-file/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectToDatabase() *sql.DB {
	conn := config.Get().DatabaseDSN
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
