package logic

import (
	"context"

	"akatm/rpc/iam/internal/svc"
	"akatm/rpc/iam/pb/iamRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UseInviteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUseInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UseInviteLogic {
	return &UseInviteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 使用邀请码
func (l *UseInviteLogic) UseInvite(in *iamRpc.UseInviteReq) (*iamRpc.UseInviteResp, error) {
	// todo: add your logic here and delete this line

	return &iamRpc.UseInviteResp{}, nil
}
