package schema

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

var WalletSchema = `
CREATE TABLE IF NOT EXISTS wallet (
	wallet_id    INTEGER PRIMARY KEY AUTO_INCREMENT,
    income INT(11)  DEFAULT 0,
    outcome  INT(11)  DEFAULT 0,
	title   VARCHAR(250) DEFAULT '',
	created_at DATETIME
);`

type Wallet struct {
	WalletId  int32  `db:"wallet_id"`
	Income    int32  `db:"income"`
	Outcome   int32  `db:"outcome"`
	Title     string `db:"title"`
	CreatedAt string `db:"created_at"`
}

type WalletRepository struct {
	Wallet *Wallet
	DB     *sqlx.DB
	wheres []Where
	limit  int32
}

func (w WalletRepository) Delete() {
	whereQuery, values := w.parseWhere()
	query := fmt.Sprintf(`DELETE FROM wallet %s`, whereQuery)

	tx := w.DB.MustBegin()
	tx.MustExec(query, values...)
	tx.Commit()
}

func (w *WalletRepository) Save() {

	wallet := w.Wallet
	income := wallet.Income
	outcome := wallet.Outcome
	title := wallet.Title
	createdAt := wallet.CreatedAt

	if createdAt == "" {
		createdAt = time.Now().Format("2006-01-02 15:04:05")
	}

	query := "INSERT INTO wallet (income, outcome, title, created_at) VALUES (?, ?, ?, ?)"

	tx := w.DB.MustBegin()
	tx.MustExec(query, income, outcome, title, createdAt)
	tx.Commit()

}

func (w *WalletRepository) Update() {

	wallet := w.Wallet
	income := wallet.Income
	outcome := wallet.Outcome
	title := wallet.Title

	whereQuery, values := w.parseWhere()
	values = append([]interface{}{income, outcome, title}, values...)
	query := `UPDATE wallet SET income=?, outcome=?, title=? ` + whereQuery

	tx := w.DB.MustBegin()
	tx.MustExec(query, values...)
	tx.Commit()
}

func (w WalletRepository) parseWhere() (string, []interface{}) {

	columns := w.wheres
	whereClause := ""
	values := []interface{}{}
	for i, column := range columns {
		if i == 0 {
			whereClause += "WHERE "
		} else {
			whereClause += column.Operator
		}

		whereClause += fmt.Sprintf("%s=?", column.Key)
		values = append(values, column.Value)
	}

	return whereClause, values
}

func (w *WalletRepository) Limit(count int32) *WalletRepository {
	w.limit = count
	return w
}

func (w WalletRepository) Get() []Wallet {

	limit := w.limit
	whereClause, values := w.parseWhere()
	query := fmt.Sprintf("SELECT * FROM wallet %s", whereClause)

	if limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}

	wallets := []Wallet{}
	w.DB.Select(&wallets, query, values...)
	return wallets
}

func (w *WalletRepository) Where(columns []Where) *WalletRepository {
	w.wheres = append(w.wheres, columns...)
	return w
}

// Check this entry is exist or not
func (w WalletRepository) Exist() bool {
	wallets := w.Limit(1).Get()
	exists := len(wallets) > 0
	if exists {
		wallet := &wallets[0]
		w.Wallet.WalletId = wallet.WalletId
		w.Wallet.Income = wallet.Income
		w.Wallet.Outcome = wallet.Outcome
		w.Wallet.Title = wallet.Title
		w.Wallet.CreatedAt = wallet.CreatedAt
	}

	return exists
}

func NewWalletRepository(db *sqlx.DB) *WalletRepository {
	return &WalletRepository{
		DB:     db,
		Wallet: &Wallet{},
	}
}
