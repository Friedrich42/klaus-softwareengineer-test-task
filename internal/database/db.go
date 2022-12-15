package database

import "github.com/jmoiron/sqlx"

type Database interface {
	Connect() error
	Connection() *sqlx.DB
}
