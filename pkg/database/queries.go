package database

import (
	"database/sql"
	"errors"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

const all_urls_query = "select * from urls"

type Database struct {
	URL string
}

func NewDatabase(url string) *Database {
	return &Database{
		URL: url,
	}
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
		return rows, err
	}
	return rows, nil
}

func (d Database) GetURLs() ([]*URL, error) {
	rows, err := d.ExecQuery(all_urls_query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var urls []*URL
	for rows.Next() {
		var url URL
		err := rows.Scan(&url.ID, &url.OriginalURL, &url.ShortenedHash, &url.DateCreated)
		if err != nil {
			return urls, err
		}

		urls = append(urls, &url)
	}

	return urls, nil
}
