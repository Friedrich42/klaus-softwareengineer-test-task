package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Sqlite struct is a helper struct for connecting to a sqlite database.
// You MUST call Connect() before using the Connection() method.
type Sqlite struct {
	conn *sqlx.DB
}

func NewSqlite() *Sqlite {
	return &Sqlite{
		conn: nil,
	}
}

func (s *Sqlite) Connect() error {
	connection, err := sqlx.Connect("sqlite3", "database.db")
	if err != nil {
		return err
	}
	s.conn = connection
	return nil
}

func (s *Sqlite) Connection() *sqlx.DB {
	return s.conn
}
