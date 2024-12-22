package config

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	tele "gopkg.in/telebot.v4"
)

func NewTelegramBot(commands []*Executor, db *sqlx.DB) {
	pref := tele.Settings{
		Token:  "7899290898:AAHeAjBgaYRaGWXdqrMqhzVp9GHOFa4UUQM",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	for _, cmd := range commands {
		b.Handle("/"+cmd.Cmd, func(c tele.Context) error {
			flag := NewCommandFlag(c.Args())
			resp := cmd.BotHandler(db, flag, c)
			return resp
		})
	}

	b.Start()
}
