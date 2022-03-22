package domain

type TransactionResponse struct {
	TransactionId   string
	AccountId       string
	NewBalance      float64
	TransactionType string
}
