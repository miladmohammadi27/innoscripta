package grpc

import (
	"context"

	"transaction/internal/api/grpc/dto"
	"transaction/internal/helper/log"
	"transaction/internal/manager"

	"github.com/samber/do"
)

type balanceServiceHandler struct {
	lg         log.Logger
	balanceMgr manager.BalanceManager
	dto.UnimplementedBalanceServiceServer
}

func NewBalanceServiceHandler(i *do.Injector) (dto.BalanceServiceServer, error) {
	return &balanceServiceHandler{
		lg:         do.MustInvoke[log.Logger](i),
		balanceMgr: do.MustInvoke[manager.BalanceManager](i),
	}, nil
}

func (th *balanceServiceHandler) UpdateBalance(ctx context.Context, req *dto.UpdateBalanceReq) (*dto.UpdateBalanceRes, error) {
	nb, err := th.balanceMgr.UpdateBalance(ctx, newEntityFromTransaction(req))
	if err != nil {
		th.lg.Error(err.Error())
		return nil, err
	}
	return &dto.UpdateBalanceRes{
		Code:       200,
		Message:    "success",
		NewBalance: nb,
	}, nil
}
