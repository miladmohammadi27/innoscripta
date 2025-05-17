package cockroach

import (
	"context"
	"time"

	"transaction/internal/helper/log"
	"transaction/internal/manager/entity"
	"transaction/internal/repo"

	"github.com/jackc/pgx/v5"
	"github.com/samber/do"
)

type balanceRepo struct {
	db *pgx.Conn
	lg log.Logger
}

func NewBalanceRepo(i *do.Injector) (repo.BalanceRepo, error) {
	return balanceRepo{
		db: do.MustInvoke[*pgx.Conn](i),
		lg: do.MustInvoke[log.Logger](i),
	}, nil
}

func (ur balanceRepo) UpdateBalance(ctx context.Context,
	accountID int32, userID int32, amount float32, trxType entity.TransactionType,
) (entity.Transaction, float32, error) {
	var accBalance accountBalance
	var trxData entity.Transaction
	err := ur.db.QueryRow(ctx,
		selectAccountBalanceQuery, accountID, userID).Scan(&accBalance.AccountID,
		&accBalance.Balance, &accBalance.UserID, &accBalance.Version)
	if err != nil {
		return trxData, 0, err
	}

	if trxType == entity.TransactionTypeWithdraw {
		if accBalance.Balance < amount {
			return trxData, accBalance.Balance, errInsufficientBalance
		}
		accBalance.Balance -= amount
	} else {
		accBalance.Balance += amount
	}

	var newBalanceRepo float32
	var updatedAt time.Time
	err = ur.db.QueryRow(ctx, updateAccountBalanceQuery,
		accBalance.Balance, accountID, userID, accBalance.Version).Scan(&newBalanceRepo, &updatedAt)
	if err != nil {
		return trxData, 0, err
	}

	trxData = entity.Transaction{
		AccountID:       accountID,
		UserID:          userID,
		Amount:          amount,
		Status:          "success",
		TransactionType: trxType,
		CreatedAt:       updatedAt,
	}

	return trxData, newBalanceRepo, nil
}
