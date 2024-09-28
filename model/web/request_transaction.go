package web

type RequestTransaction struct {
	ID     *int    `json:"id"`
	Amount float64 `json:"amount"`
	UserID int     `json:"user_id"`
	Desc   string  `json:"desc"`
}
