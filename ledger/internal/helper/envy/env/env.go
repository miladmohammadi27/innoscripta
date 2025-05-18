package env

import (
	"ledger/internal/helper/envy"

	"github.com/caarlos0/env/v7"
)

// Options is the options for the envy package, it is just a wrapper for env.Options
type Options = env.Options

type envImpl struct {
	opts Options
}

// New would create a new envy implementation with caarlos0/env/v7 with the given options
func New(opts Options) envy.Envy {
	return &envImpl{
		opts: opts,
	}
}

// Parse would take a struct and parse the environment variables into it
// it would return an error if it fails to parse the environment variables
// Options is passed during the creation of the envy package
func (e *envImpl) Parse(v interface{}) error {
	return env.Parse(v, e.opts)
}
