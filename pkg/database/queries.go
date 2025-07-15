package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

type Database struct {
	URL string
}

func (d Database) Connect() (*sql.DB, error) {
	db, err := sql.Open("libsql", d.URL)
	if err != nil {
		return nil, errors.New("Could not open connection to database.")
	}
	return db, nil
}
