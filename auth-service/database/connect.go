package database

import (
	 "github.com/namsral/flag"
)

var (
	databaseUrl = flag.String("database-url","mysql://remotemysql.com/MwKoIhSd6E?useSSL=false&autoreconnect=true","DataBase URL")
	databaseTimeout = flag.Int64("database-timeout-ms",2000,"DataBase timeout in milliseconds")
)

func Connect()  {
	// Connect to database

}
// create a database instance
func Instance() (Database,error) {

}