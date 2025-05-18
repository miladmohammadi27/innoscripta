package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"ledger/internal/helper/envy"
	"ledger/internal/helper/envy/env"
	"ledger/internal/repo"
	"ledger/internal/repo/mongo"

	helper "ledger/internal/helper/di"

	iLog "ledger/internal/helper/log"
	iNats "ledger/internal/helper/nats"

	"github.com/nats-io/nats.go"
	"github.com/samber/do"
	"github.com/urfave/cli/v2"
)

const (
	// environment variables
	appName = "LEDGER_"
)

var appCommand = cli.Command{
	Name:   "serve",
	Usage:  "Start the app",
	Action: serveApp,
}

func serveApp(cliCtx *cli.Context) error {
	// dependency injector
	injector := helper.NewInjector()

	// run shutdown on signals, listen on different goroutine
	go func() {
		if err := injector.ShutdownOnSIGTERM(); err != nil {
			log.Fatalf("failed to shutdown injector: %v", err)
		}
	}()

	// environment configurator
	do.Provide(injector, func(i *do.Injector) (envy.Envy, error) {
		return env.New(env.Options{
			Prefix: appName,
		}), nil
	})

	// logger
	do.Provide(injector, iLog.NewLogger)
	logger := do.MustInvoke[iLog.Logger](injector)
	logger.Info("Starting Ledger...")

	// queue
	do.Provide(injector, iNats.NewNatsConnection)

	// repo
	do.Provide(injector, mongo.NewLedgerRepo)

	var natsCfg iNats.NatsConfig
	if err := helper.GetConfigFromDI(injector, &natsCfg); err != nil {
		return err
	}

	nc := do.MustInvoke[*nats.Conn](injector)

	repo := do.MustInvoke[repo.LedgerRepo](injector)
	// Create a subscription
	subscription, err := nc.Subscribe(natsCfg.SubSubject, func(msg *nats.Msg) {
		err := repo.WriteLogs(msg.Data)
		if err != nil {
			logger.Error(err.Error())
		}
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe: %w", err)
	}
	defer subscription.Unsubscribe()

	// Wait for termination signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("Shutting down subscriber...")
	return nil
}
