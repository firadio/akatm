package logic

import (
	"context"

	"akatm/rpc/iam/internal/svc"
	"akatm/rpc/iam/pb/iamRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateSessionLogic {
	return &ValidateSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证会话
func (l *ValidateSessionLogic) ValidateSession(in *iamRpc.ValidateSessionReq) (*iamRpc.ValidateSessionResp, error) {
	// todo: add your logic here and delete this line

	return &iamRpc.ValidateSessionResp{}, nil
}
