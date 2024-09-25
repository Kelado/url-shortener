package repositories

import (
	"errors"

	"github.com/Kelado/url-shortener/models"
)

var (
	ErrNotFound = errors.New("not found")
)

type LinkRepository interface {
	CreateLink(*models.Link) error
	GetLink(code string) (*models.Link, error)
}
