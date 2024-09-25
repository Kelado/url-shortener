package repositories

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Kelado/url-shortener/models"
)

const (
	defaultDSN = "file:link.db?cache=shared&mode=rwc"
)

type SQLiteRepoConfig struct {
	DSN string
}

type SQLite struct {
	db *sql.DB
}

func NewSQLiteDB(conf *SQLiteRepoConfig) *SQLite {
	dsn := defaultDSN
	if conf != nil && conf.DSN != "" {
		dsn = conf.DSN
	}

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Unable to ping SQLite database ... %v", err)
	}

	err = initTables(db)
	if err != nil {
		log.Fatal(err)
	}

	return &SQLite{
		db: db,
	}
}

func (db *SQLite) CreateLink(link *models.Link) error {
	insertQuery := `INSERT INTO links (code, created_at, original_url) VALUES (?, ?, ?)`
	_, err := db.db.Exec(insertQuery, link.Code, link.CreatedAt, link.OriginalURL)
	return err
}

func (db *SQLite) GetLink(code string) (*models.Link, error) {
	row := db.db.QueryRow(`SELECT code, created_at, original_url FROM links WHERE code = ?`, code)

	var link models.Link
	err := row.Scan(&link.Code, &link.CreatedAt, &link.OriginalURL)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func initTables(db *sql.DB) error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS links (
		code TEXT PRIMARY KEY,
		created_at DATETIME,
		original_url TEXT
	);`

	if _, err := db.Exec(createTableQuery); err != nil {
		return err
	}

	return nil
}

func (db *SQLite) Drop() {
	os.Remove("./db.db")
}
