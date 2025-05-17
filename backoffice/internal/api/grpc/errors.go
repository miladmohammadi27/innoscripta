package grpc

import (
	"fmt"
)

var prefix = "HANDLR_ERROR"

var errInternalServer = fmt.Errorf("%s: internal server error", prefix)
