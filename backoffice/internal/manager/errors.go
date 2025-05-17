package manager

import "fmt"

var prefix = "MANAGER_ERROR"

var (
	ErrInvalidUser = fmt.Errorf("%s: invalid user data", prefix)
	ErrCreateUser  = fmt.Errorf("%s: error while creating user", prefix)
)

type ErrType int

const (
	INVALID_ACTION ErrType = iota
	PERSISTENCE_VIOLATION
	NON_EXISTENCE
	UNKNOWN
)

func (e ErrType) String() string {
	switch e {
	case INVALID_ACTION:
		return "invalid request"
	case PERSISTENCE_VIOLATION:
		return "unable to persist data"
	case NON_EXISTENCE:
		return "does not exits"
	default:
		return "internal error"
	}
}

type ManagerError struct {
	Type ErrType
	MSG  string
	Err  error
}

func (m ManagerError) Error() string {
	return fmt.Sprintf("%s: %s", m.Type.String(), m.MSG)
}
