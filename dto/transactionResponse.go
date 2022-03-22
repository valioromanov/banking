package dto

type TransactionResponse struct {
	TransactionId   string  `json:"transationId"`
	AccountId       string  `json:"accountId"`
	NewBalance      float64 `json:"newBalance"`
	TransactionType string  `json:"transactionType"`
	TransactionDate string  `json:"transactionDate"`
}
