package cmd

import (
	"log"

	igRPC "transaction/internal/api/grpc"
	helper "transaction/internal/helper/di"
	"transaction/internal/helper/envy"
	"transaction/internal/helper/envy/env"
	iLog "transaction/internal/helper/log"
	"transaction/internal/manager"
	"transaction/internal/repo/cockroach"
	"transaction/internal/services/nats"

	"github.com/samber/do"
	"github.com/urfave/cli/v2"
)

const (
	// environment variables
	appName = "TRANSACTION_"
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
	logger.Info("Starting transaction...")

	// database
	do.Provide(injector, cockroach.NewCockroachSingleConnection)

	// queue
	do.Provide(injector, nats.NewNatsConnection)

	// repositories
	do.Provide(injector, cockroach.NewBalanceRepo)
	do.Provide(injector, nats.NewTransactionQueue)

	// managers
	do.Provide(injector, manager.NewBalanceManager)

	// handlers
	do.Provide(injector, igRPC.NewBalanceServiceHandler)

	// grpc server
	do.Provide(injector, igRPC.NewServer)
	do.Provide(injector, igRPC.NewGateway)

	// start grpc server
	go func() {
		gServer := do.MustInvoke[igRPC.Server](injector)
		if err := gServer.Serve(); err != nil {
			logger.Error("error shutting down gRPC server")
		}
	}()

	// star grpc gateway
	gw := do.MustInvoke[igRPC.Gateway](injector)
	if err := gw.ListenAndServe(); err != nil {
		logger.Error(err.Error())
		logger.Error("shutting down gRPC gateway")
	}

	return nil
}
