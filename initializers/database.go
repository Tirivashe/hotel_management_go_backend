package initializers

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDb() {
	connString := os.Getenv("DB_STRING")
	if connString == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database!")
	DB = db
}
