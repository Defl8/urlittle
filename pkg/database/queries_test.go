package database

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func Setup() (*Database, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	url := os.Getenv("TURSO_DB_URL")
	token := os.Getenv("TURSO_DB_TOKEN")
	completeURL := url + token
	dbInst := NewDatabase(completeURL)
	return dbInst, nil
}
func TestConnect(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Errorf("Could not load .env file. ERROR: %v", err)
	}
	url := os.Getenv("TURSO_DB_URL")
	token := os.Getenv("TURSO_DB_TOKEN")
	completeURL := url + token
	dbInst := NewDatabase(completeURL)

	_, err = dbInst.Connect()
	if err != nil {
		t.Errorf("Could not connect to database with URL %s. ERROR: %v", completeURL, err)
	}
}

func TestExecuteQuery(t *testing.T) {
	dbInst, err := Setup()
	if err != nil {
		t.Errorf("Could not load .env file. ERROR: %v", err)
	}
	rows, err := dbInst.ExecQuery("SELECT id FROM urls")
	if err != nil {
		t.Errorf("Could not query the database. ERROR: %v", err)
	}
	defer rows.Close()
}

func TestGetURLs(t *testing.T) {
	expected_urls := []*URL{newURL(1, "test", "test", "2025-07-16")}

	dbInst, err := Setup()
	if err != nil {
		t.Errorf("Could not load .env file. ERROR: %v", err)
	}
	urls, err := dbInst.GetURLs()
	if err != nil {
		t.Errorf("Error getting URLs. ERROR: %v", err)
	}
	for i := range urls {
		fmt.Printf("Index %d expected URL: %+v\n", i, expected_urls[i])
		fmt.Printf("Index %d actual URL: %+v\n", i, urls[i])

		// Need to be dereferenced in order to compare by struct members.
		if *urls[i] != *expected_urls[i] {
			t.Error("URLs do not match expected")
		}
	}
}
