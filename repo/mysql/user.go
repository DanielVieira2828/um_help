package mysql

import (
	"context"

	"github.com/DanielVieirass/um_help/model"
	"github.com/jmoiron/sqlx"
)

type User struct {
	cli *sqlx.DB
}

func (r *User) InsertUser(ctx context.Context, user *model.User) error {
	query := "INSERT INTO um_help.tab_user (first_name, last_name, document_number, balance) VALUES (?, ?, ?, ?);"

	_, err := r.cli.ExecContext(ctx, query, user.FirstName, user.LastName, user.DocumentNumber, user.Balance)

	if err != nil {
		return err
	}

	return nil

}
