package logic

import (
	"context"

	"akatm/rpc/iam/internal/svc"
	"akatm/rpc/iam/pb/iamRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 用户管理相关 ============
func (l *GetUserLogic) GetUser(in *iamRpc.GetUserReq) (*iamRpc.GetUserResp, error) {
	// todo: add your logic here and delete this line

	return &iamRpc.GetUserResp{}, nil
}
