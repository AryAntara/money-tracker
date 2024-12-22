package command

import (
	"log"
	"money-tracker/config"
	"money-tracker/schema"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

func NewInsertExecutor() *config.Executor {
	return &config.Executor{
		Cmd:     "insert",
		Handler: insertHandler,
	}
}

func insertHandler(db *sqlx.DB, flag *config.CommandFlag) {

	var outcome, income int32
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
