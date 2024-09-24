package main

import "time"

type URL string

type Link struct {
	ID           string
	CreatedAt    time.Time
	ShortenedURL URL
	OriginalURL  URL
}

type LinkRequest struct {
	OriginalURL URL `json:"url"`
}
