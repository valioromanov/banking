package app

import (
	"banking/domain"
	"banking/dto"
	"banking/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AcconuntHandler struct {
	service service.AccountService
}

func (acH AcconuntHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	var reqNewAcc dto.NewAccountRequest

	err := json.NewDecoder(r.Body).Decode(&reqNewAcc)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	vars := mux.Vars(r)
	custId := vars["customer_id"]
	reqNewAcc.CustomerId = custId
	account, appError := acH.service.NewAccount(reqNewAcc)

	if appError != nil {
		writeResponse(w, appError.Code, appError)
	} else {
		writeResponse(w, http.StatusCreated, account)
	}

}

func (acH AcconuntHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	var reqMakeTrn domain.Transaction

	err := json.NewDecoder(r.Body).Decode(&reqMakeTrn)

	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	vars := mux.Vars(r)

	accountId := vars["account_id"]
	reqMakeTrn.AccountId = accountId

	trn, appError := acH.service.MakeTransaction(reqMakeTrn)

	if appError != nil {
		writeResponse(w, appError.Code, appError)
	} else {
		writeResponse(w, http.StatusCreated, trn)
	}

}

func NewAccountHandler(service service.AccountService) AcconuntHandler {
	return AcconuntHandler{service}
}
