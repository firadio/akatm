package logic

import (
	"context"

	"akatm/rpc/iam/internal/svc"
	"akatm/rpc/iam/pb/iamRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSessionLogic {
	return &CreateSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 会话管理相关 ============
func (l *CreateSessionLogic) CreateSession(in *iamRpc.CreateSessionReq) (*iamRpc.CreateSessionResp, error) {
	// todo: add your logic here and delete this line

	return &iamRpc.CreateSessionResp{}, nil
}
