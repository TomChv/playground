package database

import (
	"context"

	"encore.dev/storage/sqldb"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	// Package name is the database name
	rawDb = sqldb.Named("database").Stdlib()
	db    = sqlx.NewDb(rawDb, "postgresql")
)

// Client give access to the database client
func Client() *sqlx.DB {
	return db
}

// Placeholder used to create the SQL database
//encore:api private
func Placeholder(_ context.Context) error {
	return nil
}
