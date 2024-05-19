package model

import "time"

type User struct {
	Id             int64     `db:"user_id"`
	PublicId       string    `db:"public_id"`
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	DocumentNumber string    `db:"document_number"`
	Password       string    `db:"password"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
