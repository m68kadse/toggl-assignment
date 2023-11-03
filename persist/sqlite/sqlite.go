package sqlite

import (
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed init.sql
var initSQL string

type SQLiteDB struct {
	db            *sql.DB
	preparedStmts map[string]*sql.Stmt
}

func NewDB(filepath string) (*SQLiteDB, error) {
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

	// Create the SQLiteDB object
	sqliteDB := &SQLiteDB{
		db:            db,
		preparedStmts: make(map[string]*sql.Stmt),
	}

	// Run the SQL initialization script
	if err := sqliteDB.init(); err != nil {
		return nil, err
	}

	return sqliteDB, nil
}

func (s *SQLiteDB) PrepareStmt(query string) (*sql.Stmt, error) {
	// Check if the statement is already prepared
	if stmt, ok := s.preparedStmts[query]; ok {
		return stmt, nil
	}

	// Prepare the statement if it's not already prepared
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	// Store the prepared statement in the map
	s.preparedStmts[query] = stmt

	return stmt, nil
}

func (s *SQLiteDB) Close() error {
	// Close prepared statements
	for _, stmt := range s.preparedStmts {
		stmt.Close()
	}

	// Close the database connection
	return s.db.Close()
}

func (s *SQLiteDB) init() error {
	_, err := s.db.Exec(initSQL)
	return err
}
