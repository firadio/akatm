package system

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSystemConfigByCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取分类配置
func NewGetSystemConfigByCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSystemConfigByCategoryLogic {
	return &GetSystemConfigByCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSystemConfigByCategoryLogic) GetSystemConfigByCategory(req *types.GetSystemConfigByCategoryReq) (resp *types.GetSystemConfigByCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
