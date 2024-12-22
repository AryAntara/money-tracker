package command

import (
	"fmt"
	"log"
	"money-tracker/config"
	"money-tracker/schema"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"gopkg.in/telebot.v4"
)

func NewInsertExecutor() *config.Executor {
	return &config.Executor{
		Cmd:        "insert",
		Handler:    insertHandler,
		BotHandler: insertBotHandler,
	}
}

func insertHandler(db *sqlx.DB, flag *config.CommandFlag) {

	var inputNominal, inputTitle string

	if flag.IsPresent("-n") {
		inputNominal = flag.GetValue("-n")
	} else {
		input := createForm([]Form{
			{
				prompt: "Silakan masukan nominal nya (+/-) 	: ",
				name:   "nominal",
			},
		})
		inputNominal = input["nominal"]

	}

	if flag.IsPresent("-t") {
		inputTitle = flag.GetValue("-t")
	} else {
		input := createForm([]Form{
			{
				prompt: "Silakan masukan nama catatan 		: ",
				name:   "title",
			},
		})
		inputTitle = input["title"]
	}

	insert(db, inputNominal, inputTitle)
}

func insert(db *sqlx.DB, inputNominal string, inputTitle string) {
	var outcome, income int32
	op := strings.Split(inputNominal, "")[0]

	if op != "+" && op != "-" {
		log.Fatalf("Unknown operator \"%s\"", op)
	}

	nominal, err := strconv.Atoi(strings.ReplaceAll(inputNominal[1:], "_", ""))
	if err != nil {
		panic(err)
	}

	if op == "+" {
		income += int32(nominal)
	} else {
		outcome += int32(nominal)
	}

	walletRepository := schema.NewWalletRepository(db)
	wallet := walletRepository.Wallet
	wallet.Income = income
	wallet.Outcome = outcome
	wallet.Title = inputTitle

	walletRepository.Save()
}

func insertBotHandler(db *sqlx.DB, flag *config.CommandFlag, c telebot.Context) error {
	nominal := flag.ValueOn(0)
	operator := "-"
	if flag.ValueOn(1) == "dari" {
		operator = "+"
	}

	inputNominal := fmt.Sprintf("%s%s", operator, nominal)
	inputTitle := strings.Join(flag.Args[2:], " ")

	insert(db, inputNominal, inputTitle)
	return c.Send("Berhasil!")
}
