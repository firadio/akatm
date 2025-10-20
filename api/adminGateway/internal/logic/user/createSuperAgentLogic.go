package user

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSuperAgentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建总代
func NewCreateSuperAgentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSuperAgentLogic {
	return &CreateSuperAgentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSuperAgentLogic) CreateSuperAgent(req *types.CreateSuperAgentReq) (resp *types.CreateSuperAgentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
