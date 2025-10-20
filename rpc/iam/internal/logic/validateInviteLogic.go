package logic

import (
	"context"

	"akatm/rpc/iam/internal/svc"
	"akatm/rpc/iam/pb/iamRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateInviteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateInviteLogic {
	return &ValidateInviteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证邀请码
func (l *ValidateInviteLogic) ValidateInvite(in *iamRpc.ValidateInviteReq) (*iamRpc.ValidateInviteResp, error) {
	// todo: add your logic here and delete this line

	return &iamRpc.ValidateInviteResp{}, nil
}
