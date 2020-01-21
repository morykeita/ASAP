package database

import (
	"github.com/jmoiron/sqlx"
	"io"
)

const uniqueViolation = "unique_violation"

// Database - interface for database
type Database interface {
	UsersDB
	io.Closer
}

type database struct {
	conn *sqlx.DB
}

// close database instance  connection
func (d *database) Close() error {
	return d.conn.Close()
}


