package model

import "time"

type User struct {
	Id             int       `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	DocumentNumber string    `json:"document_number"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Deleted        time.Time `json:"deleted"`
}
