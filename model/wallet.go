package model

import "time"

type Wallet struct {
	Id          int64     `db:"wallet_id"`
	OwnerId     int64     `db:"owner_id"`
	Alias       string    `db:"alias"`
	CurrencyId  int64     `db:"curreny_id"`
	Balance     int64     `db:"balance"`
	CreatedAt   time.Time `db:"createdAt"`
	DeletedAt   time.Time `db:"deletedAt"`
	UpdatedAtAt time.Time `db:"updatedAtAt"`
}
