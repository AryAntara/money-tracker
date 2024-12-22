package command

import (
	"fmt"
	"log"
	"money-tracker/config"
	"money-tracker/helper"
	"money-tracker/schema"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"gopkg.in/telebot.v4"
)

func NewUpdateExecutor() *config.Executor {
	return &config.Executor{
		Cmd:        "update",
		Handler:    updateHandler,
		BotHandler: updateBotHandler,
	}
}

func updateHandler(db *sqlx.DB, flag *config.CommandFlag) {

	walletRepository := schema.NewWalletRepository(db)
	wallets := walletRepository.Get()

	prompt := ""
	for index, wallet := range wallets {
		prompt += fmt.Sprintf("%d. %s (-%s +%s) \n", index+1, wallet.Title,
			helper.FormatThousand(wallet.Outcome),
			helper.FormatThousand(wallet.Income))
	}

	prompt += "Pilih salah satu yang akan diganti : "
	form := createForm([]Form{{
		prompt,
		"id",
	}})

	whereClause := []schema.Where{{
		Operator: "AND",
		Key:      "wallet_id",
		Value:    form["id"],
	}}

	if !walletRepository.Where(whereClause).Exist() {
		fmt.Printf("Data %s tidak ada pada list yang ditampilkan.\n", form["id"])
		return
	}

	form = createForm([]Form{
		{
			prompt: "Masukan title baru (kosongkan bilang menggunakan yang lama) \n: ",
			name:   "title",
		},
		{
			prompt: "Masukan nominal baru (kosongkan bilang menggunakan yang lama) \n: ",
			name:   "nominal",
		},
	})

	title := form["title"]
	nominalString := form["nominal"]

	wallet := walletRepository.Wallet
	if title != "" {
		wallet.Title = title
	}

	if nominalString != "" {
		op := strings.Split(nominalString, "")[0]

		if op != "+" && op != "-" {
			log.Fatalf("Unknown operator \"%s\"", op)
		}

		nominal, err := strconv.Atoi(strings.ReplaceAll(nominalString[1:], "_", ""))
		if err != nil {
			panic(err)
		}

		if op == "+" {
			wallet.Income = int32(nominal)
			wallet.Outcome = 0
		}

		if op == "-" {
			wallet.Outcome = int32(nominal)
			wallet.Income = 0
		}
	}
	walletRepository.Update()
}

func updateBotHandler(db *sqlx.DB, flag *config.CommandFlag, c telebot.Context) error {
	return c.Send("Insert")
}
