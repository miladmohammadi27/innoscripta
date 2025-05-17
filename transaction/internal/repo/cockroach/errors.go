package cockroach

import (
	"fmt"
)

var prefix = "REPO_ERROR"

var (
	errGetCfg              = fmt.Errorf("%s: error get config", prefix)
	errInsufficientBalance = fmt.Errorf("%s: insufficient balance", prefix)
)
