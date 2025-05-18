package nats

import (
	"errors"

	"ledger/internal/helper/di"

	"github.com/nats-io/nats.go"
	"github.com/samber/do"
)

func NewNatsConnection(i *do.Injector) (*nats.Conn, error) {
	var cfg NatsConfig
	if err := di.GetConfigFromDI(i, &cfg); err != nil {
		return nil, errors.Join(errGetCfg, err)
	}

	nc, err := nats.Connect(cfg.URL)
	if err != nil {
		return nil, errors.Join(errConnection, err)
	}
	return nc, nil
}
