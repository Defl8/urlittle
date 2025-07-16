package database

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

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
