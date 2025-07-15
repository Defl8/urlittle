package database

import (
	"database/sql"
	"errors"
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

func (d Database) ExecQuery(query string, args ...any) (*sql.Rows, error) {
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(query, args...)
	if err != nil {
		return rows, errors.New("Could not execute query")
	}
	return rows, nil
}
