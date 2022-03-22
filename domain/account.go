package domain

import (
	"banking/dto"
	"banking/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponse() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	GetAccountInfo(string) (*Account, *errs.AppError)
	MakeTransaction(Transaction) (*Transaction, *errs.AppError)
}
