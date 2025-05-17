package grpc

import (
	"backoffice/internal/api/grpc/dto"
	"backoffice/internal/manager/entity"
)

func newEntityFromCreateUserReq(req *dto.CreateUserReq) entity.User {
	if req == nil {
		return entity.User{}
	}
	return entity.User{
		Name:  req.Name,
		Email: req.Email,
	}
}
