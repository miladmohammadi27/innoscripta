package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// Execute - Entrypoint for cli
func Execute() {
	app := &cli.App{Commands: []*cli.Command{
		&appCommand,
	}}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
