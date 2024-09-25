package repositories

import (
	"testing"
	"time"

	"github.com/Kelado/url-shortener/models"
	"github.com/stretchr/testify/assert"
)

func newRepository() LinkRepository {
	return NewSQLiteDB(&SQLiteRepoConfig{DSN: ":memory:"})
}

func TestCreateLink(t *testing.T) {
	repo := newRepository()

	link := models.Link{
		Code:        "FAIRLO",
		CreatedAt:   time.Now(),
		OriginalURL: "https://example.com",
	}

	err := repo.CreateLink(&link)
	assert.Nil(t, err)
}

func TestCreateLinkWithAlreadyExistingCode(t *testing.T) {
	repo := newRepository()

	link := models.Link{
		Code:        "FAIRLO",
		CreatedAt:   time.Now(),
		OriginalURL: "https://example.com",
	}

	repo.CreateLink(&link)
	err := repo.CreateLink(&link)
	assert.NotNil(t, err)
}

func TestGetExistingLink(t *testing.T) {
	repo := newRepository()
	code := "FAIRLO"
	expectedURL := models.URL("https://example.com")

	link := models.Link{
		Code:        code,
		CreatedAt:   time.Now(),
		OriginalURL: expectedURL,
	}

	repo.CreateLink(&link)
	originalLink, _ := repo.GetLink(code)
	assert.Equal(t, expectedURL, originalLink.OriginalURL)
}

func TestGetUnkonwnLink(t *testing.T) {
	repo := newRepository()
	code := "FAIRLO"
	_, err := repo.GetLink(code)
	assert.NotNil(t, err)
}
