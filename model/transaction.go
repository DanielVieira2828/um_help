package model

import "time"

type transaction struct {
	id          int       `db:"id"`
	sender_id   int       `db:"sender_id"`
	receiver_id int       `db:"receiver_id"`
	amount      int       `db:"amount"`
	createdAt   time.Time `db:"createdAt"`
}
