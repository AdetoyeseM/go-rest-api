package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "Olumide4014$"
	dbname := "my_first_db"

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error

	DB, err = sql.Open("postgres", psqlInfo) // CORRECT
	if err != nil {
		log.Fatal("Error connecting:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot ping DB:", err)
	}

	log.Println("✅ Connected to PostgreSQL successfully!")
}
