package services

import (
	"context"

	"transaction/internal/manager/entity"
)

type QueueRepo interface {
	Publish(ctx context.Context, trx entity.Transaction, newBalance float32) error
}
