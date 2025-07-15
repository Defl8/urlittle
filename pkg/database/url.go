package database

import (
	"errors"
	"fmt"
	"time"
)

type URL struct {
	ID            int
	OriginalURL   string
	ShortenedHash string
	DateCreated   time.Time
}

const timeFormat = "2006-Jan-02"

func newURL(id int, originalURL, shortenedHash, dateCreatedString string) (*URL, error) {
	dateCreated, err := time.Parse(timeFormat, dateCreatedString)
	if err != nil {
		errorString := fmt.Sprintf("'%s' could not be parsed as a date.", dateCreatedString)
		return nil, errors.New(errorString)
	}
	return &URL{
		ID:            id,
		OriginalURL:   originalURL,
		ShortenedHash: shortenedHash,
		DateCreated:   dateCreated,
	}, nil
}
