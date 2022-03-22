package app

import (
	"banking/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		writeResponse(w, err.Code, err)
		/*w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)*/
	} else {
		writeResponse(w, http.StatusOK, customers)
		/*w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)*/
	}
}

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["customer_id"]
	customers, err := ch.service.GetCustomer(id)

	if err != nil {
		writeResponse(w, err.Code, err)
		/*w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)*/
	} else {
		writeResponse(w, http.StatusOK, customers)
		/*w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)*/
	}

}

func (ch *CustomerHandlers) GetCustomerByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	status := vars["status"]
	statusInt, _ := strconv.Atoi(status)
	customers, err := ch.service.GetCustomerByStatus(statusInt)
	if err != nil {
		writeResponse(w, err.Code, err)
		/*w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)*/
	} else {
		writeResponse(w, http.StatusOK, customers)
		/*w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)*/
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
