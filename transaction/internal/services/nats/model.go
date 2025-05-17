package nats

import "transaction/internal/manager/entity"

type transaction struct {
	AccountID       int32   `json:"account_id"`
	UserID          int32   `json:"user_id"`
	Amount          float32 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	Status          string  `json:"status"`
	CreatedAt       string  `json:"created_at"`
}

func newTransactionFromEntity(entTrx entity.Transaction) transaction {
	return transaction{
		AccountID:       entTrx.AccountID,
		UserID:          entTrx.UserID,
		Amount:          entTrx.Amount,
		TransactionType: string(entTrx.TransactionType),
		Status:          entTrx.Status,
		CreatedAt:       entTrx.CreatedAt.String(),
	}
}
