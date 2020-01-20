/**
 * @author Mory Keita on 1/20/20
 */
package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/morykeita/ASAP/auth-service/internal/config"
	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	log "github.com/sirupsen/logrus"

)

func migrateDb( db *sql.DB)  error {

	driver, err := mysql.WithInstance(db,&mysql.Config{})
	if err != nil{
		return errors.Wrap(err,"connecting to database")
	}


	migrationSource := fmt.Sprintf("file://%internal/database/migrations",config.DataDirectory)
	migrator,err := migrate.NewWithDatabaseInstance(migrationSource,"mysql",driver)
	if err != nil{
		return errors.Wrap(err,"creating database migrator")
	}
	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange{
		return errors.Wrap(err,"executing migration")
	}

	version,dirty,err := migrator.Version()
	if err != nil{
		return errors.Wrap(err,"getting migration version")
	}
	log.WithFields(log.Fields{
		"version" : version,
		"dirty":dirty,
	}).Debug("Database migrated")
	return nil
}