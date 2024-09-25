package models

import (
	"errors"
	"time"
)

var (
	ErrMissingRequiredFields = errors.New("missing required field")
)

type URL string

const EmptyURL = URL("")

type Link struct {
	Code        string
	CreatedAt   time.Time
	OriginalURL URL
}

type LinkRequest struct {
	OriginalURL *URL `json:"url"`
}

func (lr *LinkRequest) ValidateSchema() error {
	if lr.OriginalURL == nil {
		return ErrMissingRequiredFields
	}
	return nil
}

type LinkResponse struct {
	ShortURL URL `json:"shortUrl"`
}
