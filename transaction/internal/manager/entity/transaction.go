package entity

import "time"

type TransactionType string

var (
	TransactionTypeWithdraw TransactionType = "WITHDRAWAL"
	TransactionTypeDeposit  TransactionType = "DEPOSIT"
)

func (t TransactionType) IsValid() bool {
	switch t {
	case TransactionTypeWithdraw, TransactionTypeDeposit:
		return true
	}
	return false
}

type Transaction struct {
	TransactionID   int32
	UserID          int32
	AccountID       int32
	Status          string
	CreatedAt       time.Time
	Amount          float32
	TransactionType TransactionType
}
