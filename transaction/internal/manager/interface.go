package manager

import (
	"context"

	"transaction/internal/manager/entity"
)

type BalanceManager interface {
	UpdateBalance(ctx context.Context, transaction entity.Transaction) (newBalance float32, err error)
}
