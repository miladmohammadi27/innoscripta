package nats

import (
	"errors"
	"log"

	"transaction/internal/helper/di"

	"github.com/nats-io/nats.go"
	"github.com/samber/do"
)

func NewNatsConnection(i *do.Injector) (nats.JetStreamContext, error) {
	var cfg NatsConfig
	if err := di.GetConfigFromDI(i, &cfg); err != nil {
		return nil, errors.Join(errGetCfg, err)
	}

	nc, err := nats.Connect(cfg.URL)
	if err != nil {
		return nil, errors.Join(errConnection, err)
	}

	// Enable JetStream
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	// Check if the stream already exists
	_, err = js.StreamInfo(cfg.PublishStream)
	if err == nats.ErrStreamNotFound {
		// If the stream doesn't exist, create it
		streamConfig := &nats.StreamConfig{
			Name:      cfg.PublishStream,
			Subjects:  []string{cfg.PublishSubject},
			Retention: nats.LimitsPolicy, // Retain based on limits
			Storage:   nats.FileStorage,  // Persistent storage
			Replicas:  1,                 // Number of replicas
		}

		_, err = js.AddStream(streamConfig)
		if err != nil {
			return nil, err
		}
		log.Printf("Stream '%s' created with subject %s.", cfg.PublishStream, cfg.PublishSubject)
	} else if err != nil {
		return nil, err
	} else {
		log.Printf("Stream '%s' already exists.", cfg.PublishStream)
	}

	return js, nil
}
