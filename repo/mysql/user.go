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

func (r *User) SelectByDocumentNumber(tx *sqlx.Tx, ctx context.Context, documentNumber string) (*model.User, bool, error) {
	var result model.User

	query := `
		SELECT
			user_id,
			public_id,
			first_name,
			last_name,
			document_number,
			password,
			created_at,
			updated_at
		FROM um_help.tab_user
		WHERE document_number = ?
		AND deleted_at IS NULL
		LIMIT 1;`

	exec := r.cli.QueryxContext
	if tx != nil {
		exec = tx.QueryxContext
	}

	rows, err := exec(ctx, query, documentNumber)
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