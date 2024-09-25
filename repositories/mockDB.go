package repositories

import "github.com/Kelado/url-shortener/models"

type MockDB struct {
	store map[string]models.Link
}

func NewMockDB() *MockDB {
	return &MockDB{
		store: make(map[string]models.Link),
	}
}

func (db *MockDB) CreateLink(link *models.Link) error {
	return nil
}

func (db *MockDB) GetLink(code string) (*models.Link, error) {
	return &models.Link{}, nil
}
