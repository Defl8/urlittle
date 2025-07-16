package database

type URL struct {
	ID            int
	OriginalURL   string
	ShortenedHash string
	DateCreated   string
}

const timeFormat = "2006-Jan-02"

func newURL(id int, originalURL, shortenedHash, dateCreated string) *URL {
	return &URL{
		ID:            id,
		OriginalURL:   originalURL,
		ShortenedHash: shortenedHash,
		DateCreated:   dateCreated,
	}
}
