package dto

type CustomerResponse struct {
	Id          string `json:"customerId"`
	Name        string `json:"fullName"`
	City        string `json:"city"`
	Zipcode     string `json:"zipCode"`
	DateofBirth string `json:"dateOfBirth"`
	Status      string `json:"status"`
}
