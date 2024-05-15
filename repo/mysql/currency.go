package mysql

import (
	"context"

	"github.com/DanielVieirass/um_help/model"
	"github.com/jmoiron/sqlx"
)

type Currency struct {
	cli *sqlx.DB
}

func (r *Currency) SelectByCurrencyCode(tx *sqlx.Tx, ctx context.Context, code model.CurrencyCode) (*model.Currency, bool, error) {
	var result model.Currency

	query := `
		SELECT
			currency_id,
			code,
			symbol,
			created_at,
			updated_at
		FROM um_help.tab_currency
		WHERE code = ?
		AND deleted_at IS NULL
		LIMIT 1;`

	exec := r.cli.QueryxContext
	if tx != nil {
		exec = tx.QueryxContext
	}

	rows, err := exec(ctx, query, code)
	if err != nil {
		return nil, false, err
	}

	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(&result); err != nil {
			return nil, false, err
		}

	} else {
		return nil, false, nil
	}

	if rows.Err() != nil {
		return nil, false, rows.Err()
	}

	return &result, true, nil
}
