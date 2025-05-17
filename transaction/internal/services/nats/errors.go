package nats

import (
	"fmt"
)

var prefix = "REPO_ERROR"

var (
	errGetCfg     = fmt.Errorf("%s: error get config", prefix)
	errConnection = fmt.Errorf("%s: error establishing nats connection", prefix)
)
