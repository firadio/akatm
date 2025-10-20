package system

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSystemConfigByCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 按分类更新系统配置
func NewUpdateSystemConfigByCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSystemConfigByCategoryLogic {
	return &UpdateSystemConfigByCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSystemConfigByCategoryLogic) UpdateSystemConfigByCategory(req *types.UpdateSystemConfigByCategoryReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
