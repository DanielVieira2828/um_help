package mysql

import (
	"context"

	"github.com/DanielVieirass/um_help/model"
	"github.com/jmoiron/sqlx"
)

type User struct {
	cli *sqlx.DB
}

func (r *User) Insert(tx *sqlx.Tx, ctx context.Context, user *model.User) (userId int64, err error) {
	query := `INSERT INTO um_help.tab_user (first_name, last_name, document_number) VALUES (?, ?, ?);`

	exec := r.cli.ExecContext
	if tx != nil {
		exec = tx.ExecContext
	}

	result, err := exec(ctx, query, user.FirstName, user.LastName, user.DocumentNumber)
	if err != nil {
		return 0, err
	}

	userId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userId, nil
}
