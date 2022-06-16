package Database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // <------------ here
	"log"
)

var db *sql.DB

func GetConnection() *sql.DB {
	return db
}

func ConnectDatabase() {
	// Capture connection properties.
	user := "postgres"
	password := "proddbpassword"
	host := "127.0.0.1"
	port := "5432"
	dbname := "shortener"

	// Get a database handle.
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("DB Error : %s", err)
	}

	fmt.Println("Database Connected!")
}
