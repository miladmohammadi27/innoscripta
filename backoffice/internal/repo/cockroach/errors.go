package cockroach

import (
	"fmt"
)

var prefix = "REPO_ERROR"

var (
	errInvalidUser = fmt.Errorf("%s: invalid user", prefix)
	errGetCfg      = fmt.Errorf("%s: error get config", prefix)
	errInsert      = fmt.Errorf("%s: error while inserting data", prefix)
)
