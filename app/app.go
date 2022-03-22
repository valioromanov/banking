package app

import (
	"banking/domain"
	"banking/service"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {
	router := mux.NewRouter()
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//wiring
	dbClient := getDbClient()
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb(dbClient))}
	ah := AcconuntHandler{service.NewAccountService(domain.NewAccountRepositoryDb(getDbClient()))}

	router.
		HandleFunc("/customers", ch.GetAllCustomers).
		Methods(http.MethodGet)
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).
		Methods(http.MethodGet)
	router.
		HandleFunc("/customers/status/{status:[0-9]+}", ch.GetCustomerByStatus).
		Methods(http.MethodGet)
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).
		Methods(http.MethodPost)
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id}", ah.MakeTransaction).
		Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:r1r2r3r4@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
