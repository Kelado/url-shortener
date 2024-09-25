package repositories

import (
	"errors"

	"github.com/Kelado/url-shortener/models"
)

type MockDB struct {
	store map[string]*models.Link
}

func NewMockDB() *MockDB {
	return &MockDB{
		store: make(map[string]*models.Link),
	}
}

func (db *MockDB) CreateLink(link *models.Link) error {
	db.store[link.Code] = link
	return nil
}

func (db *MockDB) GetLink(code string) (*models.Link, error) {
	link, ok := db.store[code]
	if !ok {
		return nil, errors.New("not found")
	}
	return link, nil
}
