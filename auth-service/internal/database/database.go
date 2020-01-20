package database

import (
	"github.com/jmoiron/sqlx"
	"io"
)
// Database - interface for database
type Database interface {
	io.Closer
}

type database struct {
	conn *sqlx.DB
}

// close database instance  connection
func (d *database) Close() error {
	return d.conn.Close()
}


