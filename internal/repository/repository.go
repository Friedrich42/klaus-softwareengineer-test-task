package repository

import (
	"cipactonal/internal/database"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	connection *sqlx.DB
}

func NewRepository(db database.Database) *Repository {
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	return &Repository{
		connection: db.Connection(),
	}
}
