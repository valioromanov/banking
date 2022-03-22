package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"banking/logger"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(domain.Transaction) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (ac DefaultAccountService) NewAccount(account dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	err := account.Validate()
	if err != nil {
		return nil, err
	}
	accountReq := domain.Account{
		AccountId:   "",
		CustomerId:  account.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      "1",
	}
	var dbResp *domain.Account
	dbResp, err = ac.repo.Save(accountReq)
	if err != nil {
		logger.Error("Error in creating new accont: " + err.Message)
		return nil, err
	}
	//response := dbResp.AccountId
	responseObj := dbResp.ToNewAccountResponse()

	return &responseObj, nil
}

func (ac DefaultAccountService) MakeTransaction(t domain.Transaction) (*dto.TransactionResponse, *errs.AppError) {

	err := t.Validate()

	if err != nil {
		return nil, err
	}

	t.TransactionDate = time.Now().Format("2006-01-02 15:04:05")

	var tr *domain.Transaction

	tr, err = ac.repo.MakeTransaction(t)

	if err != nil {
		return nil, err
	}

	response := tr.ToTransactionResponse()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
