package sqlite

import (
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed init.sql
var initSQL string

type SQLiteDAO struct {
	db            *sql.DB
	preparedStmts map[string]*sql.Stmt
}

func NewDB(filepath string) (*SQLiteDAO, error) {
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
		db:            db,
		preparedStmts: make(map[string]*sql.Stmt),
	}

	// Run the SQL initialization script
	if err := SQLiteDAO.init(); err != nil {
		return nil, err
	}

	return SQLiteDAO, nil
}

func (s *SQLiteDAO) PrepareStmt(query string) (*sql.Stmt, error) {
	// Check if the statement is already prepared
	if stmt, ok := s.preparedStmts[query]; ok {
		return stmt, nil
	}

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	s.preparedStmts[query] = stmt

	return stmt, nil
}

func (s *SQLiteDAO) Close() error {
	for _, stmt := range s.preparedStmts {
		stmt.Close()
	}

	return s.db.Close()
}

func (s *SQLiteDAO) init() error {
	_, err := s.db.Exec(initSQL)
	return err
}
