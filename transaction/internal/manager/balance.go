package manager

import (
	"context"
	"fmt"

	"transaction/internal/helper/log"
	"transaction/internal/manager/entity"
	"transaction/internal/repo"
	"transaction/internal/services"

	"github.com/samber/do"
)

type balanceManager struct {
	balanceRepo repo.BalanceRepo
	queueSrvc   services.QueueRepo
	lg          log.Logger
}

func NewBalanceManager(i *do.Injector) (BalanceManager, error) {
	return balanceManager{
		balanceRepo: do.MustInvoke[repo.BalanceRepo](i),
		queueSrvc:   do.MustInvoke[services.QueueRepo](i),
		lg:          do.MustInvoke[log.Logger](i),
	}, nil
}

func (bm balanceManager) UpdateBalance(ctx context.Context, transaction entity.Transaction) (float32, error) {
	if !transaction.TransactionType.IsValid() {
		return 0, fmt.Errorf("invalid transaction type")
	}

	trx, newBalance, err := bm.balanceRepo.UpdateBalance(ctx,
		transaction.AccountID, transaction.UserID, transaction.Amount, transaction.TransactionType)
	if err != nil {
		bm.lg.Error(err.Error())
		return 0, err
	}

	err = bm.queueSrvc.Publish(ctx, trx, newBalance)
	if err != nil {
		bm.lg.Error(err.Error())
		return 0, err
	}
	return newBalance, nil
}
