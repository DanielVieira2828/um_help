package mysql

import (
	"context"

	"github.com/DanielVieirass/um_help/model"
	"github.com/jmoiron/sqlx"
)

type User struct {
	cli *sqlx.DB
}

func (r *User) InsertUser(ctx context.Context, user *model.User) (userId int64, err error) {
	query := "INSERT INTO um_help.tab_user (first_name, last_name, document_number) VALUES (?, ?, ?);"

	result, err := r.cli.ExecContext(ctx, query, user.FirstName, user.LastName, user.DocumentNumber)

	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil

}
