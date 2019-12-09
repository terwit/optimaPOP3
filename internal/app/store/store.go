package store

import (
	"database/sql"
	
	_ "github.com/denisenkom/go-mssqldb"
	
	//_ "github.com/lib/pq"
)
// Store ...
type Store struct {
	config *Config
	db *sql.DB
}

// New ...
func New(config *Config) *Store{
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	//db, err := sql.Open("postgres",s.config.DatabaseURL)
	db, err := sql.Open("mssql",s.config.DatabaseURL)
	if err != nil{
		return err
	}
	
	if err := db.Ping(); err != nil {
		return err
	}
	
	s.db = db
	return nil
}

// Close ...
func(s *Store) Close() {
	s.db.Close()
}