package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "mydb"
)

func Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

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
