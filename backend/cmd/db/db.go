package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	if sslMode == "" {
		sslMode = "disable"
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error: Could not connect to the database")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error: Could not ping the database")
		return nil, err
	}

	fmt.Println("Successfully connected!")

	return db, nil
}
