package logic

import (
	"context"

	"akatm/rpc/iam/internal/svc"
	"akatm/rpc/iam/pb/iamRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInviteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInviteLogic {
	return &GetInviteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取邀请详情
func (l *GetInviteLogic) GetInvite(in *iamRpc.GetInviteReq) (*iamRpc.GetInviteResp, error) {
	// todo: add your logic here and delete this line

	return &iamRpc.GetInviteResp{}, nil
}
