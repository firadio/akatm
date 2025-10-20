package system

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSystemConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新系统配置
func NewUpdateSystemConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSystemConfigLogic {
	return &UpdateSystemConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSystemConfigLogic) UpdateSystemConfig(req *types.UpdateSystemConfigReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
