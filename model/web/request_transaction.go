package web

type RequestTransaction struct {
	Amount float64 `json:"amount"`
	UserID int     `json:"user_id"`
	Desc   string  `json:"desc"`
}
