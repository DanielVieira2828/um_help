package mysql

import (
	"context"

	"github.com/DanielVieirass/um_help/model"
	"github.com/jmoiron/sqlx"
)

type Wallet struct {
	cli *sqlx.DB
}

func (r *Wallet) InsertWallet(ctx context.Context, wallet *model.Wallet) error {
	query := "INSERT INTO um_help.tab_wallet (owner_id, alias, currency_id) VALUES (?, ?, ?);"

	_, err := r.cli.ExecContext(ctx, query, wallet.OwnerId, wallet.Alias, wallet.CurrencyId)

	if err != nil {
		return err
	}

	return nil

}
