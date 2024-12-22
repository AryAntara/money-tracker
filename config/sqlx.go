package config

import (
	"log"
	"money-tracker/schema"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewDatabase() Database {
	driverName := "sqlite3"
	dbName := "source.db"

	db, err := sqlx.Connect(driverName, dbName)
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
