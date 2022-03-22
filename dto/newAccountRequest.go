package dto

import "banking/errs"

type NewAccountRequest struct {
	CustomerId  string  `json:"customerId"`
	AccountId   string  `json:"accountId"`
	Amount      float64 `json:"amount"`
	AccountType string  `json:"accountType"`
}

func (na NewAccountRequest) Validate() *errs.AppError {
	if na.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit at least 5000")
	}

	if na.AccountType != "saving" && na.AccountType != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
