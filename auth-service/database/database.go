package database

import (
	"github.com/jmoiron/sqlx"
	"io"
)
// nameDatabase - interface for database
type Database interface {
	io.Closer
}

type database struct {
	conn *sqlx.DB
}
// close database instance  connection
func (d* database) close() error {
	return d.conn.Close()
}