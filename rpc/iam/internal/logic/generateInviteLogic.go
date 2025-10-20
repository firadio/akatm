package logic

import (
	"context"

	"akatm/rpc/iam/internal/svc"
	"akatm/rpc/iam/pb/iamRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateInviteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateInviteLogic {
	return &GenerateInviteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 邀请管理相关 ============
func (l *GenerateInviteLogic) GenerateInvite(in *iamRpc.GenerateInviteReq) (*iamRpc.GenerateInviteResp, error) {
	// todo: add your logic here and delete this line

	return &iamRpc.GenerateInviteResp{}, nil
}
