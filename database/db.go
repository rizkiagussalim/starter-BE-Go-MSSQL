package database

import (
	"database/sql"
	"fmt"
	"log"
	// "os"

	_ "github.com/denisenkom/go-mssqldb"
)

func InitDB(dbUser, dbPass, dbHost, dbName string) (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", dbHost, dbUser, dbPass, dbName)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
