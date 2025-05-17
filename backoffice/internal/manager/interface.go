package manager

import (
	"context"

	"backoffice/internal/manager/entity"
)

type UserManager interface {
	CreateUser(ctx context.Context, user entity.User) (userID string, err error)
	CreateAccount(ctx context.Context, userID string) (accountID int32, err error)
}
