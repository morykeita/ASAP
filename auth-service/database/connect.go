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
	databaseUrl = flag.String("database-url","GG9MDssj6z:hu0iDf1Ogj@tcp(remotemysql.com)/GG9MDssj6z","DataBase URL")
	databaseTimeout = flag.Int64("database-timeout-ms",200000,"DataBase timeout in milliseconds")
	max_connection = 32
)
//user:password@/dbname"


func Connect()(*sqlx.DB, error) {
	// Connect to database
	dbUrl := *databaseUrl
	log.Debug("connecting to database")
	conn, err := sqlx.Open("mysql", dbUrl)
	if err != nil {
		return nil , errors.Wrap(err,"could not connect to database")
	}
	conn.SetMaxOpenConns(max_connection)

	// check if database is running
	if err := waitForDB(conn.DB) ; err != nil {
		return nil , err
	}
	return conn,nil
}

// create a database instance
func New() (Database,error) {
	conn, err := Connect()
	if err != nil {
		return  nil,err
	}
	d := &database{conn:conn}
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