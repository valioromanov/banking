package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Valentin", "Rakovski", "4150", "1996-08-29", "1"},
		{"1002", "Ivan", "Plovdiv", "4000", "1996-01-29", "1"},
	}

	return CustomerRepositoryStub{customers}
}
