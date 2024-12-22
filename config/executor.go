package config

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/telebot.v4"
)

type Executor struct {
	Cmd        string
	Handler    func(db *sqlx.DB, flag *CommandFlag)
	BotHandler func(db *sqlx.DB, flag *CommandFlag, ctx telebot.Context) error
}
