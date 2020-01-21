package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/namsral/flag"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	// use postgres
	// disable sslmode
	databaseUrl     = flag.String("database-url","postgres://postgres:postgres@localhost/postgres?sslmode=disable","DataBase URL.")
	databaseTimeout = flag.Int64("database-timeout-ms",20000,"DataBase timeout in milliseconds.")
	maxConnection   = 32
)
// Connect creates a new database connection
func Connect()(*sqlx.DB, error) {
	// Connect to database
	dbUrl := *databaseUrl
	log.WithField("url",dbUrl).Debug("connecting to database.")
	conn, err := sqlx.Open("postgres", dbUrl)
	if err != nil {
		return nil , errors.Wrap(err,"could not connect to database.")
	}
	conn.SetMaxOpenConns(maxConnection)

	// check if database is running
	if err := waitForDB(conn.DB) ; err != nil {
		return nil , err
	}

	// Migrate databases schemas

	 if err := migrateDb(conn.DB); err != nil{
		return nil, errors.Wrap(err, "Could not migrate database.")
	}
	return conn,nil
}

// create a database instance
func New() (Database,error) {
	conn, err := Connect()
	if err != nil {
		return  nil,err
	}
	d := &database{conn: conn}
	return  d,nil
}

func waitForDB(conn *sql.DB) error {
	ready := make(chan struct{})
	go func() {
		for  {
			if err := conn.Ping(); err != nil{
				close(ready)
				return
			}
			time.Sleep(100*time.Millisecond)
		}
	}()

	select {
	case <- ready:
		return nil
	case <- time.After(time.Duration(*databaseTimeout)*time.Millisecond) :
		return errors.New("database not ready")
	}
}