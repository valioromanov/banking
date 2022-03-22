package domain

import (
	"banking/dto"
	"banking/errs"
	"banking/logger"
	"strings"
)

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) IsDeposit() bool {
	return strings.ToUpper(t.TransactionType) == "DEPOSIT"
}

func (t Transaction) Validate() *errs.AppError {
	if t.Amount < 0 {
		logger.Error("Negative Amount")
		return errs.NewValidationError("Negative Amount")
	}

	if strings.ToUpper(t.TransactionType) != "DEPOSIT" && strings.ToUpper(t.TransactionType) != "WITHDRAWAL" {
		logger.Error("Unsupported trasaction type")
		return errs.NewValidationError("Unsupported transaction type")
	}

	return nil
}

func (t Transaction) ValidateAmount(accBalance float64) *errs.AppError {
	if !t.IsDeposit() && t.Amount > accBalance {
		logger.Error("Amount of the transaction is bigger than the balance of the acoount")
		return errs.NewValidationError("Insufficient availability")
	}

	return nil
}

func (t Transaction) ToTransactionResponse() dto.TransactionResponse {

	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		NewBalance:      t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
