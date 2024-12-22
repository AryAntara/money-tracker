package command

import (
	"fmt"
	"money-tracker/config"

	"github.com/jmoiron/sqlx"
)

func NewResetExecutor() *config.Executor {
	return &config.Executor{
		Cmd: "reset",
		Handler: func(db *sqlx.DB, flag *config.CommandFlag) {
			db.MustExec("drop table wallet")
			fmt.Println("Database was dropped!")
		},
	}
}
