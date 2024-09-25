package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Kelado/url-shortener/internal/random"
	"github.com/Kelado/url-shortener/models"
	"github.com/Kelado/url-shortener/repositories"
)

var (
	ErrEmptyURL     = fmt.Errorf("empty url")
	ErrURLNotExists = fmt.Errorf("url does not exist")
)

type Controller struct {
	codeSize int
	hostname string
	linkRepo repositories.LinkRepository
}

func NewController(hostname string, codeSize int, linkRepo repositories.LinkRepository) *Controller {
	return &Controller{
		codeSize: codeSize,
		hostname: hostname,
		linkRepo: linkRepo,
	}
}

func (c *Controller) CreateLink(linkReq models.LinkRequest) (models.URL, error) {
	code := random.NewString(c.codeSize)

	link := models.Link{
		Code:        code,
		CreatedAt:   time.Now(),
		OriginalURL: linkReq.OriginalURL,
	}

	if err := ValidateLink(&link); err != nil {
		return models.EmptyURL, err
	}

	err := c.linkRepo.CreateLink(&link)
	if err != nil {
		return models.EmptyURL, err
	}

	return c.createShortURL(link.Code), nil
}

func (c *Controller) GetLink(code string) (models.URL, error) {
	link, err := c.linkRepo.GetLink(code)
	if err != nil {
		return models.EmptyURL, err
	}
	return link.OriginalURL, nil
}

func ValidateLink(l *models.Link) error {
	if l.OriginalURL == models.EmptyURL {
		return ErrEmptyURL
	}

	_, err := http.Head(string(l.OriginalURL))
	if err != nil {
		return ErrURLNotExists
	}

	return nil
}

func (c *Controller) createShortURL(code string) models.URL {
	return models.URL(c.hostname + code)
}
