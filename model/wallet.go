package model

import "time"

type Wallet struct {
	Id         int64     `db:"wallet_id"`
	OwnerId    int64     `db:"owner_id"`
	CurrencyId int64     `db:"currency_id"`
	Alias      string    `db:"alias"`
	Balance    int64     `db:"balance"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
