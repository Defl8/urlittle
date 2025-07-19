package database

type URL struct {
	ID            int
	OriginalURL   string
	ShortenedHash string
	DateCreated   string
}

const timeFormat = "2006-Jan-02"

func newURL(originalURL, shortenedHash, dateCreated string) *URL {
	return &URL{
		OriginalURL:   originalURL,
		ShortenedHash: shortenedHash,
		DateCreated:   dateCreated,
	}
}
