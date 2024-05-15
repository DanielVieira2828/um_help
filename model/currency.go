package model

import "time"

type Currency struct {
	Id        int64        `db:"currency_id"`
	Code      CurrencyCode `db:"code"`
	Symbol    string       `db:"symbol"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}

type CurrencyCode string

const (
	CurrencyBRL CurrencyCode = "BRL"
	CurrencyEUR CurrencyCode = "EUR"
	CurrencyUSD CurrencyCode = "USD"
)
