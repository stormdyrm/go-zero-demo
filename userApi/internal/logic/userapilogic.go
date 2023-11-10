package logic

import (
	"context"

	"userApi/internal/svc"
	"userApi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserApiLogic {
	return &UserApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserApiLogic) UserApi(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
