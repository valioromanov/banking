package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (a AccountRepositoryDb) Save(ac Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts(customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"

	result, err := a.client.Exec(sqlInsert, ac.CustomerId, ac.OpeningDate, ac.AccountType, ac.Amount, ac.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewAccountCreatingError("Cannot create account")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting new account id: " + err.Error())
		return nil, errs.NewAccountCreatingError("Cannot create account")
	}

	ac.AccountId = strconv.FormatInt(id, 10)
	return &ac, nil
}

func (a AccountRepositoryDb) GetAccountInfo(acId string) (*Account, *errs.AppError) {
	sqlGetAccInfo := "SELECT * FROM accounts where account_id = ?"

	var acInf Account
	err := a.client.Get(&acInf, sqlGetAccInfo, acId)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while searching account " + acId + ": " + sql.ErrNoRows.Error())
			return nil, errs.NewNotFoundError("Account not found")
		} else {
			logger.Error("Unexpected error: " + err.Error())
			return nil, errs.NewUnexceptedError(err.Error())
		}
	}
	return &acInf, nil
}

func (a AccountRepositoryDb) MakeTransaction(t Transaction) (*Transaction, *errs.AppError) {

	account, errCustom := a.GetAccountInfo(t.AccountId)

	tx, err := a.client.Begin()

	if err != nil {
		logger.Error("Error while beging a transaction: " + err.Error())
		return nil, errs.NewUnexceptedError(err.Error())
	}
	if errCustom != nil {
		return nil, errCustom
	}

	errCustom = t.ValidateAmount(account.Amount)
	if errCustom != nil {
		return nil, errCustom
	}

	result, _ := tx.Exec("INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?,?,?,?)",
		t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.IsDeposit() {
		_, err = tx.Exec("UPDATE accounts set AMOUNT = AMOUNT + ? WHERE account_id = ?", t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec("UPDATE accounts set AMOUNT = AMOUNT - ? WHERE account_id = ?", t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
	}

	err = tx.Commit()

	if err != nil {
		logger.Error("Error while comminting a transaction: " + err.Error())
		return nil, errs.NewUnexceptedError(err.Error())
	}

	transactionId, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting a transaction id: " + err.Error())
		return nil, errs.NewUnexceptedError(err.Error())
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)

	t.Amount = account.Amount
	return &t, nil

}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
