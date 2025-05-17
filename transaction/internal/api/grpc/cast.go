package grpc

import (
	"transaction/internal/api/grpc/dto"
	"transaction/internal/manager/entity"
)

func newEntityFromTransaction(transaction *dto.UpdateBalanceReq) entity.Transaction {
	if transaction == nil {
		return entity.Transaction{}
	}
	return entity.Transaction{
		Amount:          transaction.Amount,
		UserID:          transaction.UserID,
		AccountID:       transaction.AccountID,
		TransactionType: entity.TransactionType(transaction.TransactionType.String()),
	}
}
