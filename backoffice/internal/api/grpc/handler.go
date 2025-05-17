package grpc

import (
	"context"

	"backoffice/internal/api/grpc/dto"
	"backoffice/internal/helper/log"
	"backoffice/internal/manager"

	"github.com/samber/do"
)

type userServiceHandler struct {
	lg     log.Logger
	usrMgr manager.UserManager
	dto.UnimplementedUserServiceServer
}

func NewUserServiceHandler(i *do.Injector) (dto.UserServiceServer, error) {
	return &userServiceHandler{
		lg:     do.MustInvoke[log.Logger](i),
		usrMgr: do.MustInvoke[manager.UserManager](i),
	}, nil
}

func (us *userServiceHandler) CreateUser(ctx context.Context, req *dto.CreateUserReq) (*dto.CreateUserRes, error) {
	uID, err := us.usrMgr.CreateUser(ctx, newEntityFromCreateUserReq(req))
	if err != nil {
		return nil, err
	}

	return &dto.CreateUserRes{
		Code:    1,
		Message: "success",
		UserID:  uID,
	}, nil
}

func (us *userServiceHandler) CreateAccount(ctx context.Context, req *dto.CreateAccountReq) (*dto.CreateAccountRes, error) {
	aID, err := us.usrMgr.CreateAccount(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return &dto.CreateAccountRes{
		Code:      1,
		Message:   "success",
		AccountID: aID,
	}, nil
}
