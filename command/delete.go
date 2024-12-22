package command

import (
	"fmt"
	"money-tracker/config"
	"money-tracker/helper"
	"money-tracker/schema"

	"github.com/jmoiron/sqlx"
)

func NewDeleteExecutor() *config.Executor {
	return &config.Executor{
		Cmd:     "delete",
		Handler: deleteHandler,
	}
}

func deleteHandler(db *sqlx.DB, flag *config.CommandFlag) {
	walletRepository := schema.NewWalletRepository(db)
	wallets := walletRepository.Get()

	prompt := ""
	for index, wallet := range wallets {
		prompt += fmt.Sprintf("%d. %s (-%s +%s) \n", index+1, wallet.Title,
			helper.FormatThousand(wallet.Outcome),
			helper.FormatThousand(wallet.Income))
	}

	prompt += "Pilih salah satu yang akan dihapus : "
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

	confirmation := createForm([]Form{
		{
			prompt: "Yakin untuk menghapus data ini ? ",
			name:   "confirm",
		},
	})["confirm"]

	if confirmation == "y" {
		walletRepository.Delete()
	} else {
		fmt.Println("Tidak jadi.")
	}
}
