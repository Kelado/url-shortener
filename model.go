package main

import "time"

type URL string

type Link struct {
	Code        string
	CreatedAt   time.Time
	OriginalURL URL
}

type LinkRequest struct {
	OriginalURL URL `json:"url"`
}
