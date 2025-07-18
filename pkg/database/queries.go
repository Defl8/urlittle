package database

import (
	"database/sql"
	"errors"
	"fmt"

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
	fmt.Println("Attempting to connect to the database...")
	db, err := sql.Open("libsql", d.URL)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to the database.")
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
	fmt.Println("Fetching URLs from the database...")
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
	fmt.Printf("Returned %d URLs from the database.\n", len(urls))
	return urls, nil
}

func (d Database) AddURL(url *URL) error {
	fmt.Println("Attempting to insert a new URL into the database.")
	_, err := d.ExecQuery("insert into urls(original_url, shortened_has, date_created) values(?, ?, ?)", url.OriginalURL, url.ShortenedHash, url.DateCreated)
	if err != nil {
		return err
	}
	fmt.Println("URL was successfully inserted into the database.")
	return nil
}
