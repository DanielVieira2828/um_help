package mysql

import (
	"context"

	"github.com/DanielVieirass/um_help/model"
	"github.com/jmoiron/sqlx"
)

type Wallet struct {
	cli *sqlx.DB
}

func (r *Wallet) Insert(tx *sqlx.Tx, ctx context.Context, wallet *model.Wallet) error {
	query := `INSERT INTO um_help.tab_wallet(owner_id, currency_id, alias) VALUES (?, ?, ?);`

	exec := r.cli.ExecContext
	if tx != nil {
		exec = tx.ExecContext
	}

	_, err := exec(ctx, query, wallet.OwnerId, wallet.CurrencyId, wallet.Alias)
	if err != nil {
		return err
	}

	return nil
}
