package web

import "time"

type ResponseTransaction struct {
	Amount    float64   `json:"amount"`
	Desc      string    `json:"desc"`
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
