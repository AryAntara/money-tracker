package command

import (
	"fmt"
	"money-tracker/config"
	"money-tracker/helper"
	"money-tracker/schema"
	"time"

	"github.com/jmoiron/sqlx"
)

func NewSeeExecutor() *config.Executor {
	return &config.Executor{
		Cmd:     "see",
		Handler: SeeHandler,
	}
}

func SeeHandler(db *sqlx.DB, flag *config.CommandFlag) {
	walletRepository := schema.NewWalletRepository(db)
	wallets := walletRepository.Get()

	var incomeTotal int32 = 0
	var outcomeTotal int32 = 0
	for i, wallet := range wallets {
		i++
		date, _ := time.Parse(time.RFC3339, wallet.CreatedAt)
		fmt.Printf("%d. (%s) %s +%s -%s\n", i, date.Format("2006-01-02"), wallet.Title,
			helper.FormatThousand(wallet.Income),
			helper.FormatThousand(wallet.Outcome))

		incomeTotal += wallet.Income
		outcomeTotal += wallet.Outcome
	}

	fmt.Printf("Total Income \t: %s\nTotal Outcome\t: %s\nTotal\t\t: %s\n",
		helper.FormatThousand(incomeTotal),
		helper.FormatThousand(outcomeTotal),
		helper.FormatThousand((incomeTotal - outcomeTotal)),
	)

}
