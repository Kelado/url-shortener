package models

import (
	"time"
)

type URL string

const EmptyURL = URL("")

type Link struct {
	Code        string
	CreatedAt   time.Time
	OriginalURL URL
}

type LinkRequest struct {
	OriginalURL URL `json:"url"`
}

type LinkResponse struct {
	ShortURL URL `json:"shortUrl"`
}
