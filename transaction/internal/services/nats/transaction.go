package nats

import (
	"context"
	"encoding/json"
	"errors"

	"transaction/internal/helper/di"
	"transaction/internal/helper/log"
	"transaction/internal/manager/entity"
	"transaction/internal/services"

	"github.com/nats-io/nats.go"
	"github.com/samber/do"
)

type transactionQueue struct {
	js  nats.JetStreamContext
	lg  log.Logger
	cfg NatsConfig
}

func NewTransactionQueue(i *do.Injector) (services.QueueRepo, error) {
	var cfg NatsConfig
	if err := di.GetConfigFromDI(i, &cfg); err != nil {
		return nil, errors.Join(errGetCfg, err)
	}

	js := do.MustInvoke[nats.JetStreamContext](i)
	return &transactionQueue{
		js:  js,
		lg:  do.MustInvoke[log.Logger](i),
		cfg: cfg,
	}, nil
}

func (q transactionQueue) Publish(ctx context.Context, entTrx entity.Transaction, newBalance float32) error {
	trx, err := json.Marshal(newTransactionFromEntity(entTrx))
	if err != nil {
		return err
	}

	_, err = q.js.Publish(q.cfg.PublishSubject, []byte(trx))
	if err != nil {
		return err
	}

	return nil
}
