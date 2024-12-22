package config

import "github.com/jmoiron/sqlx"

type Executor struct {
	Cmd     string
	Handler func(db *sqlx.DB, flag *CommandFlag)
}
