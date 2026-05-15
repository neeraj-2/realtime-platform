package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectPostgres() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=admin password=password dbname=realtime sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error connecting to PostgreSQL: %v", err)
		return nil, err
	}
	
	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging PostgreSQL: %v", err)
		return nil, err
	}
	
	log.Println("Successfully connected to PostgreSQL")
	return db, nil
}