package web

type RequestTransaction struct {
	Amount float64 `json:"amount"`
	Desc   string  `json:"desc"`
}
