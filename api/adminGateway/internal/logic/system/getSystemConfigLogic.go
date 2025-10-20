package system

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSystemConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取系统配置
func NewGetSystemConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSystemConfigLogic {
	return &GetSystemConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSystemConfigLogic) GetSystemConfig(req *types.GetSystemConfigReq) (resp *types.GetSystemConfigResp, err error) {
	// todo: add your logic here and delete this line

	return
}
