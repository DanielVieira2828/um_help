package model

import "time"

type User struct {
	Id             int64     `db:"user_id"`
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	DocumentNumber string    `db:"document_number"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
