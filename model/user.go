package model

import "time"

type User struct {
	id             int       `json:"id"`
	FirstName      string    `json:"name"`
	LastName       string    `json:"last_name"`
	DocumentNumber string    `json:"document_number"`
	Balance        int       `json:"balance"`
	createdAt      time.Time `json:"createdAt"`
	updatedAt      time.Time `json:"updatedAt"`
	deleted        time.Time `json:"deleted"`
}
