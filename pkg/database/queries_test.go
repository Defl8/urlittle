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
	defer rows.Close()
	if err != nil {
		t.Errorf("Could not query the database. ERROR: %v", err)
	}
}
