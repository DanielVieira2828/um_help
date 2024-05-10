package model

import "time"

type transaction struct {
	id          int       `json:"id"`
	sender_id   int       `json:"sender_id"`
	receiver_id int       `json:"receiver_id"`
	amount      int       `json:"amount"`
	createdAt   time.Time `json:"createdAt"`
}
