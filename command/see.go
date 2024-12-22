package command

import (
	"fmt"
	"money-tracker/config"
	"money-tracker/helper"
	"money-tracker/schema"
	"time"

	"github.com/jmoiron/sqlx"
	"gopkg.in/telebot.v4"
)

func NewSeeExecutor() *config.Executor {
	return &config.Executor{
		Cmd:        "see",
		Handler:    SeeHandler,
		BotHandler: seeBotHandler,
	}
}

func generateView(db *sqlx.DB, flag *config.CommandFlag) string {
	walletRepository := schema.NewWalletRepository(db)
	wallets := walletRepository.Get()

	var incomeTotal int32 = 0
	var outcomeTotal int32 = 0
	var daily string

	for i, wallet := range wallets {
		i++
		date, _ := time.Parse(time.DateTime, wallet.CreatedAt)

		daily += fmt.Sprintf("%d. (%s) %s +%s -%s\n", i, date.Format("2006-01-02"), wallet.Title,
			helper.FormatThousand(wallet.Income),
			helper.FormatThousand(wallet.Outcome))

		incomeTotal += wallet.Income
		outcomeTotal += wallet.Outcome
	}

	return daily + fmt.Sprintf("Total Income \t: %s\nTotal Outcome\t: %s\nTotal\t\t: %s\n",
		helper.FormatThousand(incomeTotal),
		helper.FormatThousand(outcomeTotal),
		helper.FormatThousand((incomeTotal-outcomeTotal)),
	)

}
func SeeHandler(db *sqlx.DB, flag *config.CommandFlag) {
	fmt.Println(generateView(db, flag))
}

func seeBotHandler(db *sqlx.DB, flag *config.CommandFlag, c telebot.Context) error {
	return c.Send(fmt.Sprintf("```\n%s```", generateView(db, flag)), &telebot.SendOptions{
		ParseMode: telebot.ModeMarkdown,
	})
}
