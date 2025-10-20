package managerProfile

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFeeSettingsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新费率设置
func NewUpdateFeeSettingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFeeSettingsLogic {
	return &UpdateFeeSettingsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFeeSettingsLogic) UpdateFeeSettings(req *types.UpdateFeeSettingsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
