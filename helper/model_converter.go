package helper

import (
	"github.com/rizky201008/mywallet-backend/model/domain"
	"github.com/rizky201008/mywallet-backend/model/web"
)

func TransactionToResponseTransaction(data domain.Transaction) web.ResponseTransaction {
	var desc string
	if data.Desc != nil {
		desc = *data.Desc
	} else {
		desc = ""
	}
	return web.ResponseTransaction{
		CreatedAt: data.CreatedAt,
		Amount:    data.Amount,
		ID:        int(data.ID),
		Desc:      desc,
	}
}

func TransactionsToResponseTransactions(data []domain.Transaction) []web.ResponseTransaction {
	var response []web.ResponseTransaction
	for _, value := range data {
		var desc string
		if value.Desc != nil {
			desc = *value.Desc
		} else {
			desc = ""
		}

		response = append(response, web.ResponseTransaction{
			CreatedAt: value.CreatedAt,
			Amount:    value.Amount,
			ID:        int(value.ID),
			Desc:      desc,
		})
	}

	return response
}

func RequestTransactionToTransaction(data web.RequestTransaction) domain.Transaction {
	return domain.Transaction{
		Amount: data.Amount,
		Desc:   &data.Desc,
		UserID: data.UserID,
	}
}
