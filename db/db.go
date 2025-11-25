package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func Connect() {
	host := ""
	port := 
	user := ""
	password := ""
	dbname := ""

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot ping DB:", err)
	}

	log.Println("✅ Connected to PostgreSQL successfully!")
}
