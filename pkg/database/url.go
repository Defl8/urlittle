package database

import "time"

type URL struct {
	ID            int
	OriginalURL   string
	ShortenedHash string
	DateCreated   time.Time
}
