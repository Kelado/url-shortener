package repositories

import (
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
	_, exists := db.store[link.Code]
	if exists {
		return ErrCodeAlreadyExists
	}
	db.store[link.Code] = link
	return nil
}

func (db *MockDB) GetLink(code string) (*models.Link, error) {
	link, ok := db.store[code]
	if !ok {
		return nil, ErrNotFound
	}
	return link, nil
}
