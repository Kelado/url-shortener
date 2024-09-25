package repositories

import (
	"github.com/Kelado/url-shortener/models"
)

type LinkRepository interface {
	CreateLink(*models.Link) error
	GetLink(code string) (*models.Link, error)
}
