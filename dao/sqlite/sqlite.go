package sqlite

import (
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed init.sql
var initSQL string

type SQLiteDAO struct {
	db *sql.DB
}

func NewDAO(filepath string) (*SQLiteDAO, error) {
	// Open the SQLite database
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Create the SQLiteDAO object
	SQLiteDAO := &SQLiteDAO{
		db: db,
	}

	// Run the SQL initialization script
	if err := SQLiteDAO.init(); err != nil {
		return nil, err
	}

	return SQLiteDAO, nil
}

func (s *SQLiteDAO) init() error {
	_, err := s.db.Exec(initSQL)
	return err
}
