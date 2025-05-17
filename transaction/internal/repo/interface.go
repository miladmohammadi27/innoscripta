package repo

import (
	"context"

	"transaction/internal/manager/entity"
)

type BalanceRepo interface {
	UpdateBalance(ctx context.Context, accountID int32, userID int32, amount float32, trxType entity.TransactionType) (entity.Transaction, float32, error)
}
