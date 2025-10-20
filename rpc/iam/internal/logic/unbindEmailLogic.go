package logic

import (
	"context"

	"akatm/rpc/iam/internal/svc"
	"akatm/rpc/iam/pb/iamRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnbindEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnbindEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindEmailLogic {
	return &UnbindEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 解绑邮箱
func (l *UnbindEmailLogic) UnbindEmail(in *iamRpc.UnbindEmailReq) (*iamRpc.UnbindEmailResp, error) {
	// todo: add your logic here and delete this line

	return &iamRpc.UnbindEmailResp{}, nil
}
