package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/ymolodyukov/otus/internal/db"
)

var _ db.DB = &Person{}

// Person реализация db.Person
type Person struct {
	DB *sqlx.DB
}
