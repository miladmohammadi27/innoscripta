package repo

import (
	"context"

	"backoffice/internal/manager/entity"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user entity.User) (string, error)
	CreateAccount(ctx context.Context, userID int) (int32, error)
}
