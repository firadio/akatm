package system

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSystemStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取系统状态
func NewGetSystemStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSystemStatusLogic {
	return &GetSystemStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSystemStatusLogic) GetSystemStatus() (resp *types.GetSystemStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
