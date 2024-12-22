package command

import (
	"fmt"
	"money-tracker/config"
	"money-tracker/helper"
	"money-tracker/schema"
	"strings"

	"github.com/jmoiron/sqlx"
	"gopkg.in/telebot.v4"
)

type analyticItem struct {
	title string
	value string
}

func NewAnalyticExecutor() *config.Executor {
	return &config.Executor{
		Cmd:        "analytic",
		Handler:    analyticHandler,
		BotHandler: analyticBotHandler,
	}
}

func analyticHandler(db *sqlx.DB, flag *config.CommandFlag) {
	walletRepository := schema.NewWalletRepository(db)
	wallets := walletRepository.Get()
	analyzed := analyze(wallets)
	analyticItems := []analyticItem{
		{
			title: "Pengeluaran terbesar harian\t",
			value: fmt.Sprintf("%s untuk %s", analyzed["most_outcome_nominal_in_day"], analyzed["most_outcome_title_in_day"]),
		}, {
			title: "Total Belanja online\t\t",
			value: fmt.Sprintf("%s (%s)", analyzed["total_buy_from_online"], analyzed["total_nominal_buy_from_online"]),
		},
	}

	for _, analyticItem := range analyticItems {
		fmt.Println(analyticItem.title, " : ", analyticItem.value)
	}
}

func analyze(wallets []schema.Wallet) map[string]string {

	var mostOutcomeNominalInDay, totalNominalBuyFromOnline int32
	mostOutcomeTitleInDay := ""
	totalBuyFromOnline := 0
	for _, wallet := range wallets {
		title := wallet.Title
		// income := wallet.Income
		outcome := wallet.Outcome
		if mostOutcomeNominalInDay < outcome {
			mostOutcomeNominalInDay = outcome
			mostOutcomeTitleInDay = title
		}

		titleComponents := strings.Split(title, " ")
		if strings.ToUpper(titleComponents[0]) == "BELI" && strings.ToUpper(titleComponents[1]) == "ONLINE" {
			totalBuyFromOnline += 1
			totalNominalBuyFromOnline += outcome
		}
	}

	return map[string]string{
		"most_outcome_nominal_in_day":   helper.FormatThousand(mostOutcomeNominalInDay),
		"most_outcome_title_in_day":     mostOutcomeTitleInDay,
		"total_buy_from_online":         fmt.Sprintf("%d Unit", totalBuyFromOnline),
		"total_nominal_buy_from_online": helper.FormatThousand(totalNominalBuyFromOnline),
	}
}

func analyticBotHandler(db *sqlx.DB, flag *config.CommandFlag, c telebot.Context) error {
	return c.Send("Insert")
}
