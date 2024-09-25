package controller

import (
	"time"

	"github.com/Kelado/url-shortener/internal/random"
	"github.com/Kelado/url-shortener/models"
	"github.com/Kelado/url-shortener/repositories"
)

const (
	CodeSize = 6

	Hostname = "http://localhost:8000/"
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

	return c.createServiceURL(link.Code), nil
}

func (c *Controller) GetLink(shortenedURL string) (models.URL, error) {
	return "originalURL", nil
}

func (c *Controller) createServiceURL(code string) models.URL {
	return models.URL(c.hostname + code)
}
