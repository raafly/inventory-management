package config

import (
	"database/sql"

	"github.com/raafly/inventory-management/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	helper.PanicIfError(err)

	return db
}