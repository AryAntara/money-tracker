package config

import (
	"log"
	"money-tracker/schema"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDatabase() Database {
	driverName := "mysql"
	// dbName := "source.db"
	dsn := os.Getenv("DSN")
	db, err := sqlx.Connect(driverName, dsn)
	if err != nil {
		log.Fatalf("Cannot connect into database : %v", err)
	}

	return Database{
		Sqlx: db,
	}
}

type Database struct {
	Sqlx *sqlx.DB
}

func (d *Database) Migrate() {
	d.Sqlx.MustExec(schema.WalletSchema)
}
