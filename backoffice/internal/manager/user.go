package manager

import (
	"context"
	"errors"
	"strconv"

	"backoffice/internal/helper/log"
	"backoffice/internal/manager/entity"
	"backoffice/internal/repo"

	"github.com/samber/do"
)

type userManager struct {
	usrRepo repo.UserRepo
	lg      log.Logger
}

func NewUserManager(i *do.Injector) (UserManager, error) {
	return userManager{
		usrRepo: do.MustInvoke[repo.UserRepo](i),
		lg:      do.MustInvoke[log.Logger](i),
	}, nil
}

func (um userManager) CreateUser(ctx context.Context, user entity.User) (string, error) {
	err := user.Validate()
	if err != nil {
		return "", errors.Join(ErrInvalidUser, err)
	}

	id, err := um.usrRepo.CreateUser(ctx, user)
	if err != nil {
		return id, errors.Join(ErrCreateUser, err)
	}
	return id, nil
}

func (um userManager) CreateAccount(ctx context.Context, userID string) (int32, error) {
	usrID, err := strconv.Atoi(userID)
	if err != nil {
		return 0, errors.Join(ErrInvalidUser, err)
	}

	uID, err := um.usrRepo.CreateAccount(ctx, usrID)
	if err != nil {
		return uID, errors.Join(ErrCreateUser, err)
	}
	return uID, nil
}
