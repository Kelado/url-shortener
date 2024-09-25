package repositories

import (
	"errors"

	"github.com/Kelado/url-shortener/models"
)

var (
	ErrNotFound          = errors.New("not found")
	ErrCodeAlreadyExists = errors.New("code already exists")
)

type LinkRepository interface {
	CreateLink(*models.Link) error
	GetLink(code string) (*models.Link, error)
}
