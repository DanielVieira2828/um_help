package res

import "time"

type User struct {
	ID        int64     `json:"userId"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}
